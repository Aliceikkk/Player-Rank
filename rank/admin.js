async function handleLogin(event) {
    event.preventDefault();
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    showLoading('正在登录...');
    
    try {
        const response = await fetch(`${API_CONFIG.BASE_URL}/api/admin/login`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                username: username,
                password: password
            })
        });

        hideLoading();

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const data = await response.json();
        if (data.success) {
            showToast('登录成功！', 'success');
            setTimeout(() => {
                document.getElementById('loginForm').style.display = 'none';
                document.getElementById('adminPanel').style.display = 'block';
                loadConfig();
            }, 500);
        } else {
            showAlert(data.message || '用户名或密码错误', 'error', '登录失败');
        }
    } catch (error) {
        hideLoading();
        console.error('Login error:', error);
        showAlert('网络连接失败，请检查后端服务是否正常运行', 'error', '登录失败');
    }
}

async function loadConfig() {
    try {
        const response = await fetch(`${API_CONFIG.BASE_URL}/api/admin/config`);
        const config = await response.json();
        
        document.getElementById('announcement').value = config.announcement || '';
        document.getElementById('lastSaveTime').value = (config.last_save_time || '').replace(' ', 'T');
        document.getElementById('syncInterval').value = config.sync_interval_minutes || 1;
    } catch (error) {
        showToast('加载配置失败: ' + error.message, 'error');
    }
}

async function saveConfig() {
    let lastSaveTime = document.getElementById('lastSaveTime').value.replace('T', ' ');
    if (lastSaveTime.length === 16) {
        lastSaveTime += ":00";
    }

    const config = {
        announcement: document.getElementById('announcement').value,
        last_save_time: lastSaveTime,
        sync_interval_minutes: parseInt(document.getElementById('syncInterval').value)
    };

    showLoading('正在保存配置...');

    try {
        const response = await fetch(`${API_CONFIG.BASE_URL}/api/admin/config`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(config)
        });

        hideLoading();

        const data = await response.json();
        if (data.success) {
            showToast('配置保存成功！', 'success');
        } else {
            showAlert(data.message || '保存失败', 'error', '保存失败');
        }
    } catch (error) {
        hideLoading();
        showAlert('网络错误: ' + error.message, 'error', '保存失败');
    }
}

async function clearLeaderboard() {
    const confirmed = await showConfirm(
        '确定要清空排行榜吗？<br><strong style="color: #f44336;">此操作不可恢复！</strong>',
        '危险操作',
        {
            confirmText: '确认清空',
            cancelText: '取消',
            type: 'error'
        }
    );
    
    if (!confirmed) {
        return;
    }

    showLoading('正在清空排行榜...');

    try {
        const response = await fetch(`${API_CONFIG.BASE_URL}/api/admin/clear`, {
            method: 'POST'
        });

        hideLoading();

        const data = await response.json();
        if (data.success) {
            showToast('排行榜已清空！', 'success');
        } else {
            showAlert(data.message || '清空失败', 'error', '操作失败');
        }
    } catch (error) {
        hideLoading();
        showAlert('网络错误: ' + error.message, 'error', '操作失败');
    }
} 