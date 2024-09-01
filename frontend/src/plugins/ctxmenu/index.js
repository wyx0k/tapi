export default {
    install(app) {
        app.directive('no-ctx-menu', {
            mounted(el) {
                el.addEventListener('contextmenu', (e) => {
                    e.preventDefault();
                });
            },
        });
    },
};