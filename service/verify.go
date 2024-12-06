package service

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"data/logger"
)

const (
	privateKeyPEM = `
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAz+MSD/1AnDAQcmXiRry0c3G1OJOBDCuwPulnzTtM1eAteoEK
NqQmY+k/8MXgtdZXwGIdS3LyHksnheMKiiSRWrcQCPcXxt0F13lksMs+Q6Uhkm+K
+fdMSfZYd5rHhWT9DQYOErCSG4aIC7KXNFAxJdcR2rF1A4ZQZOm/ALIxHyqBBHX1
7IDPoMSxp216TGC8oNEJuyXnaOiRI9Zi7ICJ76aEbOyYyFHQQz8QVxNqnfwMMR4R
4OjNaoYPVeAV49g416wB+U+900gkAOxogvbj65dW5ubYgKGzTzsqNY5fZ8EoOEP+
ptaZD1Nb9T+t4xbpVQRKnldLktfQ8wjUdPt0+wIDAQABAoIBACt0OHtVIaNEPLvc
5h6JJWMYOpw34PqtSyrrWQJCSf5O3NLFlF/1kfzCbdYjVqRLyI3zNFJoFYoVhY5r
9fEIUcXXjCeDltm/q4/t1QlLgplbNBhK7o7vjWfMcZOWq+wLPwXw5RItDG1qgpeH
XinWNV/Gg/nlrCLveXCOl5lKpg/wTvWVFTcQZEr9u168pHe68JkDWBaWC0oon4Gn
4mVC0knpzDzgWVMx18KrxvHszE7VuAkw8tZ7Opcbfayksh0cCWQO6D2e9QEJJMSs
iQMY0PXN11QV+1Hag+j/xvL4tOnU4FMLnnx4+Atrs3/wF/bUP00CbMG6z2XaRJlk
Tv6o5bECgYEA24ALJaAqy2WGTC+G8sVoCVTx+32AvE+2uUSx+F6d9ZHcpehN0ESU
Cq+9ZG17P0YWX30g2KUBUGDVCH9GDD9kvqYOXeoLj7zMkSU+Wv7++P91WjokGO4/
5/Cm7IsINJqoBwBobW1816fjyeKaPB1CBiqat6GPe+byV50a1niH7KMCgYEA8nSp
s2xa4yW75ROV9rrD/tnNziAUP8gcfVzm29EBqk946YMrm8Q/VUE3Jyl94+01smeJ
9WPufrHb2AYK1wWHobwabix4n8hdwIilS1Jpku4ZQenTX2PtkAbr/5J/l2JXJpB0
xq/u1aw15whuitpT7CsSU52ARene0BIYhIScQ8kCgYEAqIVqCX0q3fYYYa2rogBF
m8SH9Fmq2OkqlJtVOCcMh1lxWj88XsYIExxSACS+FxosWyuqaCpnE9sEM/3jPcv+
ARFkvl3OepCtTlKVyS81et5Grvssc6eXkO+GeN1Vc225y4ZYposE6l2P2ZQMblLY
OfvDxXBYxPvO281WqYrocJkCgYEAonq6CuqPUe/EpjRk6C6rEAvIFt16lG+3X++1
KhQ7yHVvsxVbUQzH4ItOuajdm7QoqwEl+9PXTQJGMNODyisDhMh48eJNh86PvWoV
U4/L/lIfjQN0ylU+K6nn16LGbJ3Th52BFHFsXbcFfGkzI0xCt1757hz7jF65Gbk2
wv/XdLkCgYAhI2RhcrKkm+N+6GEE7pZ6VN+ne4VbVK1TboV3m5c7qSeAHJogm/rn
gcNjuNCZNASe5CVt6bniHKC9Iz/Y0P73m9ME+ANksa4ZXGG1u1ewlWvgB5VDmOpj
UywPU74tcaw06SI6zqt/mF58B3YvSgbXLUxmJTPaXvs490M/nf+sxA==
-----END RSA PRIVATE KEY-----
`
)

type VerifyService struct {
	privateKey *rsa.PrivateKey
}

func NewVerifyService() *VerifyService {
	vs := &VerifyService{}
	vs.initKeys()
	return vs
}

func (vs *VerifyService) initKeys() {
	// 解析私钥
	block, _ := pem.Decode([]byte(privateKeyPEM))
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		logger.Error("解析私钥失败:", err)
		return
	}
	vs.privateKey = privateKey
}

// 生成机器唯一ID
func (vs *VerifyService) GenerateMachineID() string {
	var identifiers []string

	// 获取CPU信息
	if cpuInfo, err := ioutil.ReadFile("/proc/cpuinfo"); err == nil {
		// 只获取CPU型号
		lines := strings.Split(string(cpuInfo), "\n")
		for _, line := range lines {
			if strings.Contains(line, "model name") {
				cpuModel := strings.TrimSpace(strings.Split(line, ":")[1])
				identifiers = append(identifiers, cpuModel)
				break
			}
		}
	}

	// 获取IP地址
	if ips, err := vs.getIPAddresses(); err == nil && len(ips) > 0 {
		// 只使用第一个IP地址
		identifiers = append(identifiers, ips[0])
	}

	// 组合并生成唯一ID
	combinedString := strings.Join(identifiers, "|")
	// Base64编码
	return base64.StdEncoding.EncodeToString([]byte(combinedString))
}

// 验证激活码
func (vs *VerifyService) VerifyActivationCode(code string, machineID string) bool {
	decodedCode, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return false
	}

	// 使用私钥解密
	decrypted, err := rsa.DecryptPKCS1v15(nil, vs.privateKey, decodedCode)
	if err != nil {
		return false
	}

	// 验证解密后的内容是否与机器ID匹配
	return string(decrypted) == machineID
}

// 检查激活状态
func (vs *VerifyService) CheckActivation() bool {
	// 查激活文件是否存在
	content, err := ioutil.ReadFile("actcode.txt")
	if err != nil {
		return false
	}

	// 取当前机器ID
	machineID := vs.GenerateMachineID()

	// 验证激活码
	return vs.VerifyActivationCode(string(content), machineID)
}

// 保存激活码
func (vs *VerifyService) SaveActivationCode(code string) error {
	return ioutil.WriteFile("actcode.txt", []byte(code), 0644)
}

// 获取硬盘序列号
func (vs *VerifyService) getDiskSerial() (string, error) {
	if runtime.GOOS == "windows" {
		return vs.getWindowsDiskSerial()
	} else {
		return vs.getLinuxDiskSerial()
	}
}

// Windows系统获取硬盘序列号
func (vs *VerifyService) getWindowsDiskSerial() (string, error) {
	cmd := exec.Command("wmic", "diskdrive", "get", "serialnumber")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// 解析输出内容
	lines := strings.Split(string(output), "\n")
	if len(lines) >= 2 {
		// 第二行包含序列号
		serial := strings.TrimSpace(lines[1])
		if serial != "" {
			return serial, nil
		}
	}

	return "", fmt.Errorf("无法获取硬盘序列号")
}

// Linux系统获取硬盘序列号
func (vs *VerifyService) getLinuxDiskSerial() (string, error) {
	// 尝试读取主要磁盘设备的序列号
	devices := []string{"/dev/sda", "/dev/nvme0n1", "/dev/hda"}

	for _, device := range devices {
		// 使用hdparm命令获序列号
		cmd := exec.Command("hdparm", "-i", device)
		output, err := cmd.Output()
		if err == nil {
			// 解析输出查找SerialNo
			outputStr := string(output)
			if strings.Contains(outputStr, "SerialNo=") {
				parts := strings.Split(outputStr, "SerialNo=")
				if len(parts) > 1 {
					serial := strings.Split(parts[1], " ")[0]
					return strings.TrimSpace(serial), nil
				}
			}
		}

		// 尝试从/sys/block路径读取
		deviceName := filepath.Base(device)
		serialPath := fmt.Sprintf("/sys/block/%s/serial", deviceName)
		if serial, err := ioutil.ReadFile(serialPath); err == nil {
			return strings.TrimSpace(string(serial)), nil
		}
	}

	// 如果上述方法都失败，尝试使用udevadm命令
	cmd := exec.Command("udevadm", "info", "--query=property", "--name=/dev/sda")
	output, err := cmd.Output()
	if err == nil {
		lines := strings.Split(string(output), "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "ID_SERIAL_SHORT=") {
				return strings.TrimPrefix(line, "ID_SERIAL_SHORT="), nil
			}
		}
	}

	return "", fmt.Errorf("无法获取硬盘序列号")
}

// 获取IP地址列表
func (vs *VerifyService) getIPAddresses() ([]string, error) {
	var ips []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}
	return ips, nil
}
