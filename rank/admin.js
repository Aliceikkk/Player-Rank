async function handleLogin(event) {
    event.preventDefault();
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

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

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const data = await response.json();
        if (data.success) {
            document.getElementById('loginForm').style.display = 'none';
            document.getElementById('adminPanel').style.display = 'block';
            loadConfig();
        } else {
            alert('登录失败: ' + data.message);
        }
    } catch (error) {
        console.error('Login error:', error);
        alert('登录失败: ' + error.message);
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
        alert('加载配置失败: ' + error.message);
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

    try {
        const response = await fetch(`${API_CONFIG.BASE_URL}/api/admin/config`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(config)
        });

        const data = await response.json();
        if (data.success) {
            alert('配置保存成功');
        } else {
            alert('保存失败: ' + data.message);
        }
    } catch (error) {
        alert('保存失败: ' + error.message);
    }
}

async function clearLeaderboard() {
    if (!confirm('确定要清空排行榜吗？此操作不可恢复！')) {
        return;
    }

    try {
        const response = await fetch(`${API_CONFIG.BASE_URL}/api/admin/clear`, {
            method: 'POST'
        });

        const data = await response.json();
        if (data.success) {
            alert('排行榜已清空');
        } else {
            alert('清空失败: ' + data.message);
        }
    } catch (error) {
        alert('清空失败: ' + error.message);
    }
} 