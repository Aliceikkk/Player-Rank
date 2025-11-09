/* ==================== 自定义弹窗系统 ==================== */

// 创建弹窗容器（如果不存在）
function createModalContainer() {
    if (!document.getElementById('modal-overlay')) {
        const overlay = document.createElement('div');
        overlay.id = 'modal-overlay';
        overlay.className = 'modal-overlay';
        document.body.appendChild(overlay);
        
        // 点击遮罩层关闭弹窗
        overlay.addEventListener('click', function(e) {
            if (e.target === overlay) {
                closeModal();
            }
        });
    }
    
    if (!document.getElementById('toast-container')) {
        const toastContainer = document.createElement('div');
        toastContainer.id = 'toast-container';
        toastContainer.className = 'toast-container';
        document.body.appendChild(toastContainer);
    }
}

// Alert 弹窗
function showAlert(message, type = 'info', title = null) {
    return new Promise((resolve) => {
        createModalContainer();
        
        const overlay = document.getElementById('modal-overlay');
        
        // 设置默认标题
        if (!title) {
            const titles = {
                'success': '成功',
                'error': '错误',
                'warning': '警告',
                'info': '提示'
            };
            title = titles[type] || '提示';
        }
        
        // 设置图标
        const icons = {
            'success': '✓',
            'error': '✕',
            'warning': '!',
            'info': 'i'
        };
        
        overlay.innerHTML = `
            <div class="modal-container">
                <div class="modal-header">
                    <h3 class="modal-title">
                        <span class="modal-icon ${type}">${icons[type]}</span>
                        ${title}
                    </h3>
                    <button class="modal-close" onclick="closeModal()">×</button>
                </div>
                <div class="modal-body">
                    ${message}
                </div>
                <div class="modal-footer">
                    <button class="modal-btn modal-btn-primary" onclick="closeModal(); window.modalResolve();">
                        确定
                    </button>
                </div>
            </div>
        `;
        
        overlay.classList.add('active');
        
        // 保存 resolve 函数到全局，供按钮调用
        window.modalResolve = resolve;
        
        // ESC 键关闭
        const escHandler = (e) => {
            if (e.key === 'Escape') {
                closeModal();
                resolve();
                document.removeEventListener('keydown', escHandler);
            }
        };
        document.addEventListener('keydown', escHandler);
    });
}

// Confirm 确认弹窗
function showConfirm(message, title = '确认', options = {}) {
    return new Promise((resolve) => {
        createModalContainer();
        
        const overlay = document.getElementById('modal-overlay');
        
        const confirmText = options.confirmText || '确定';
        const cancelText = options.cancelText || '取消';
        const type = options.type || 'warning';
        
        const icons = {
            'success': '✓',
            'error': '✕',
            'warning': '!',
            'info': 'i'
        };
        
        overlay.innerHTML = `
            <div class="modal-container">
                <div class="modal-header">
                    <h3 class="modal-title">
                        <span class="modal-icon ${type}">${icons[type]}</span>
                        ${title}
                    </h3>
                    <button class="modal-close" onclick="closeModal(); window.modalResolve(false);">×</button>
                </div>
                <div class="modal-body">
                    ${message}
                </div>
                <div class="modal-footer">
                    <button class="modal-btn modal-btn-secondary" onclick="closeModal(); window.modalResolve(false);">
                        ${cancelText}
                    </button>
                    <button class="modal-btn ${type === 'error' ? 'modal-btn-danger' : 'modal-btn-primary'}" 
                            onclick="closeModal(); window.modalResolve(true);">
                        ${confirmText}
                    </button>
                </div>
            </div>
        `;
        
        overlay.classList.add('active');
        
        window.modalResolve = resolve;
        
        // ESC 键取消
        const escHandler = (e) => {
            if (e.key === 'Escape') {
                closeModal();
                resolve(false);
                document.removeEventListener('keydown', escHandler);
            }
        };
        document.addEventListener('keydown', escHandler);
    });
}

// 关闭弹窗
function closeModal() {
    const overlay = document.getElementById('modal-overlay');
    if (overlay) {
        overlay.classList.remove('active');
    }
}

// Toast 提示
function showToast(message, type = 'info', duration = 3000) {
    createModalContainer();
    
    const container = document.getElementById('toast-container');
    const toast = document.createElement('div');
    toast.className = `toast ${type}`;
    
    const icons = {
        'success': '✓',
        'error': '✕',
        'warning': '!',
        'info': 'i'
    };
    
    toast.innerHTML = `
        <div class="toast-icon">${icons[type]}</div>
        <div class="toast-content">${message}</div>
        <button class="toast-close" onclick="this.parentElement.remove()">×</button>
    `;
    
    container.appendChild(toast);
    
    // 自动关闭
    if (duration > 0) {
        setTimeout(() => {
            toast.style.animation = 'slideOutRight 0.3s ease';
            setTimeout(() => {
                if (toast.parentElement) {
                    toast.remove();
                }
            }, 300);
        }, duration);
    }
    
    return toast;
}

// 加载提示
function showLoading(message = '加载中...') {
    createModalContainer();
    
    const overlay = document.getElementById('modal-overlay');
    
    overlay.innerHTML = `
        <div class="modal-container">
            <div class="modal-body" style="text-align: center; padding: 40px;">
                <div style="margin-bottom: 20px;">
                    <div style="display: inline-block; width: 50px; height: 50px; border: 4px solid #f3f3f3; border-top: 4px solid #4CAF50; border-radius: 50%; animation: spin 1s linear infinite;"></div>
                </div>
                <div style="color: #666; font-size: 14px;">${message}</div>
            </div>
        </div>
    `;
    
    // 添加旋转动画
    if (!document.getElementById('loading-animation-style')) {
        const style = document.createElement('style');
        style.id = 'loading-animation-style';
        style.textContent = `
            @keyframes spin {
                0% { transform: rotate(0deg); }
                100% { transform: rotate(360deg); }
            }
        `;
        document.head.appendChild(style);
    }
    
    overlay.classList.add('active');
}

// 关闭加载提示
function hideLoading() {
    closeModal();
}

// 替换原生 alert 和 confirm（可选）
window.customAlert = showAlert;
window.customConfirm = showConfirm;
window.customToast = showToast;
window.customLoading = showLoading;
window.customHideLoading = hideLoading;
