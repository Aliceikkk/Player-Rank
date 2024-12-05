// 配置部分
const API_BASE_URL = 'http://192.168.1.81:8080';
const API_KEY = 'genshinimpact2023';  // 修改为与后端config.json中的api_secret一致

// 生成签名函数
async function generateSignature(params) {
    const timestamp = Math.floor(Date.now() / 1000);
    
    // 按字母顺序排序参数
    const sortedParams = Object.keys(params).sort().reduce((acc, key) => {
        acc[key] = params[key];
        return acc;
    }, {});

    // 构建签名字符串
    let signStr = Object.entries(sortedParams)
        .map(([key, value]) => `${key}=${value}`)
        .join('&');
        
    // 如果有参数，添加 & 
    if (signStr) {
        signStr += '&';
    }
    
    // 添加时间戳和密钥
    signStr += `timestamp=${timestamp}&key=${API_KEY}`;
    
    //console.log('签名字符串:', signStr);
    
    // 纯 JavaScript 实现的 SHA-256
    function sha256(ascii) {
        function rightRotate(value, amount) {
            return (value >>> amount) | (value << (32 - amount));
        }
        
        const mathPow = Math.pow;
        const maxWord = mathPow(2, 32);
        const lengthProperty = 'length';
        const words = [];
        const asciiBitLength = ascii[lengthProperty] * 8;
        
        let hash = sha256.h = sha256.h || [];
        const k = sha256.k = sha256.k || [];
        let primeCounter = k[lengthProperty];
        let result = '';

        const isComposite = {};
        for (let candidate = 2; primeCounter < 64; candidate++) {
            if (!isComposite[candidate]) {
                for (let i = 0; i < 313; i += candidate) {
                    isComposite[i] = candidate;
                }
                hash[primeCounter] = (mathPow(candidate, .5) * maxWord) | 0;
                k[primeCounter++] = (mathPow(candidate, 1 / 3) * maxWord) | 0;
            }
        }
        
        ascii += '\x80';
        while (ascii[lengthProperty] % 64 - 56) ascii += '\x00';
        for (let i = 0; i < ascii[lengthProperty]; i++) {
            let j = ascii.charCodeAt(i);
            if (j >> 8) return;
            words[i >> 2] |= j << ((3 - i) % 4) * 8;
        }
        words[words[lengthProperty]] = ((asciiBitLength / maxWord) | 0);
        words[words[lengthProperty]] = (asciiBitLength);
        
        for (let j = 0; j < words[lengthProperty];) {
            const w = words.slice(j, j += 16);
            let oldHash = hash;
            hash = hash.slice(0, 8);
            
            for (let i = 0; i < 64; i++) {
                const i2 = i + j;
                const w15 = w[i - 15], w2 = w[i - 2];
                
                const a = hash[0], e = hash[4];
                const temp1 = hash[7]
                    + (rightRotate(e, 6) ^ rightRotate(e, 11) ^ rightRotate(e, 25))
                    + ((e & hash[5]) ^ ((~e) & hash[6]))
                    + k[i]
                    + (w[i] = (i < 16) ? w[i] : (
                        w[i - 16]
                        + (rightRotate(w15, 7) ^ rightRotate(w15, 18) ^ (w15 >>> 3))
                        + w[i - 7]
                        + (rightRotate(w2, 17) ^ rightRotate(w2, 19) ^ (w2 >>> 10))
                    ) | 0);
                const temp2 = (rightRotate(a, 2) ^ rightRotate(a, 13) ^ rightRotate(a, 22))
                    + ((a & hash[1]) ^ (a & hash[2]) ^ (hash[1] & hash[2]));
                
                hash = [(temp1 + temp2) | 0].concat(hash);
                hash[4] = (hash[4] + temp1) | 0;
            }
            
            for (let i = 0; i < 8; i++) {
                hash[i] = (hash[i] + oldHash[i]) | 0;
            }
        }
        
        for (let i = 0; i < 8; i++) {
            for (let j = 3; j + 1; j--) {
                const b = (hash[i] >> (j * 8)) & 255;
                result += ((b < 16) ? 0 : '') + b.toString(16);
            }
        }
        return result;
    }

    try {
        const signature = sha256(signStr);
        //console.log('生成的签名:', signature);
        
        return {
            signature,
            timestamp
        };
    } catch (error) {
        console.error('生成签名时出错:', error);
        throw error;
    }
}

// 页面逻辑部分
let currentPage = {
    damage: 1,
    flight: 1
};

let currentTab = 'damage';
let lastUpdateTimestamp = 0;

function switchTab(tab) {
    // 移除有active类
    document.querySelectorAll('.tab-btn').forEach(btn => btn.classList.remove('active'));
    document.querySelectorAll('.leaderboard').forEach(board => board.classList.remove('active'));

    // 添加active类到选中的标签
    document.querySelector(`button[onclick="switchTab('${tab}')"]`).classList.add('active');
    document.getElementById(`${tab}-board`).classList.add('active');

    currentTab = tab;
    // 切换标签时重置为第一页
    currentPage[tab] = 1;
    updateLeaderboards();
}

function formatNumber(num) {
    return num.toLocaleString('zh-CN', { maximumFractionDigits: 2 });
}

function createPagination(type, totalPages, currentPage) {
    const pagination = document.getElementById(`${type}-pagination`);
    pagination.innerHTML = '';

    // 上一页按钮
    const prevBtn = document.createElement('button');
    prevBtn.className = 'page-btn';
    prevBtn.textContent = '上一页';
    prevBtn.disabled = currentPage === 1;
    prevBtn.onclick = () => changePage(type, currentPage - 1);
    pagination.appendChild(prevBtn);

    // 页码按钮
    for (let i = 1; i <= totalPages; i++) {
        const pageBtn = document.createElement('button');
        pageBtn.className = `page-btn ${i === currentPage ? 'active' : ''}`;
        pageBtn.textContent = i;
        pageBtn.onclick = () => changePage(type, i);
        pagination.appendChild(pageBtn);
    }

    // 下一页按钮
    const nextBtn = document.createElement('button');
    nextBtn.className = 'page-btn';
    nextBtn.textContent = '下一页';
    nextBtn.disabled = currentPage === totalPages;
    nextBtn.onclick = () => changePage(type, currentPage + 1);
    pagination.appendChild(nextBtn);
}

function changePage(type, page) {
    currentPage[type] = page;
    updateLeaderboards();
}

async function updateLeaderboards() {
    const params = {
        page: currentPage[currentTab]
    };

    const { signature, timestamp } = await generateSignature(params);
    const apiUrl = `${API_BASE_URL}/api/leaderboards?page=${params.page}&timestamp=${timestamp}&signature=${signature}`;
    
    fetch(apiUrl)
        .then(response => response.json())
        .then(data => {
            // 更新最强一击排行榜
            const damageList = document.querySelector('#damage-board .rank-list');
            const damageData = data.damage.data;
            damageList.innerHTML = damageData.map((player, index) => `
                <div class="rank-item">
                    <div class="rank rank-${index + 1}">${(data.damage.page - 1) * 8 + index + 1}</div>
                    <div class="player-info">
                        <div class="player-name">${player.nickname}</div>
                        <div class="player-score">${formatNumber(player.max_critical_damage)}</div>
                    </div>
                </div>
            `).join('');

            // 更新最远飞行距离排行榜
            const flightList = document.querySelector('#flight-board .rank-list');
            const flightData = data.flight.data;
            flightList.innerHTML = flightData.map((player, index) => `
                <div class="rank-item">
                    <div class="rank rank-${index + 1}">${(data.flight.page - 1) * 8 + index + 1}</div>
                    <div class="player-info">
                        <div class="player-name">${player.nickname}</div>
                        <div class="player-score">${formatNumber(player.max_fly_distance)}米</div>
                    </div>
                </div>
            `).join('');

            // 更新分页
            createPagination('damage', data.damage.totalPages, data.damage.page);
            createPagination('flight', data.flight.totalPages, data.flight.page);

            // 更新间戳
            if (data.timestamp) {
                lastUpdateTimestamp = data.timestamp;
            }
        })
        .catch(error => console.error('Error fetching leaderboards:', error));
}

async function updateAnnouncement() {
    try {
        const { signature, timestamp } = await generateSignature({});
        const apiUrl = `${API_BASE_URL}/api/announcement?timestamp=${timestamp}&signature=${signature}`;
        
        const response = await fetch(apiUrl, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        });
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        
        const data = await response.json();
        //console.log('公告数据:', data);
        
        if (data && data.announcement) {
            document.getElementById('announcement-content').textContent = data.announcement;
        }
    } catch (error) {
        console.error('获取公告出错:', error);
    }
}

async function checkAndRefresh() {
    const refreshBtn = document.querySelector('.refresh-btn');
    refreshBtn.classList.add('refreshing');

    const params = {
        lastUpdate: lastUpdateTimestamp
    };

    const { signature, timestamp } = await generateSignature(params);
    const apiUrl = `${API_BASE_URL}/api/check-update?lastUpdate=${lastUpdateTimestamp}&timestamp=${timestamp}&signature=${signature}`;

    fetch(apiUrl)
        .then(response => response.json())
        .then(data => {
            if (data.hasUpdate) {
                lastUpdateTimestamp = data.timestamp;
                updateLeaderboards();
                updateAnnouncement();
            }
        })
        .catch(error => console.error('Error checking updates:', error))
        .finally(() => {
            refreshBtn.classList.remove('refreshing');
        });
}
// 初始加载
updateLeaderboards();
updateAnnouncement();

