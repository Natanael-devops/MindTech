
    var header = document.getElementById('header');
    var barraNavegacao = document.getElementById('barra');
    var conteudo = document.getElementById('conteudo');
    var aparecerSideBar = false;

    
    function mudarSideBar() {
        aparecerSideBar = !aparecerSideBar;
        if (aparecerSideBar) {
            barraNavegacao.style.marginLeft = '-10vw';
            barraNavegacao.style.animationName = 'aparecerSideBar';
            conteudo.style.filter = 'blur(2px)';
        } else {
            barraNavegacao.style.marginLeft = '-100vw';
            barraNavegacao.style.animationName = '';
            conteudo.style.filter = "";
        }
    }

    function fecharSideBar() {
        if (aparecerSideBar) {
            mudarSideBar();
        }
    }

    window.addEventListener('resize', function(event) {
        if (window.innerWidth > 768 && aparecerSideBar) {
            mudarSideBar();
        }
    });


