(function ($) {

    const userLang = navigator.language || navigator.userLanguage;
    if (!userLang.includes('cs') && window.location.href.includes('mfLeaderBoard.html')) {
        if (userLang.includes('de')) {
            const ok = confirm('It seems that you may prefer German language over Czech.\nDo you want to switch to German language?\nOK = German\nCancel = Čeština');
            if (ok) window.location.href = './index_de.html';
        } else {
            const ok = confirm('It seems that you may prefer English language over Czech.\nDo you want to switch to English language?\nOK = English\nCancel = Čeština');
            if (ok) window.location.href = './index_en.html';
        }
    }

    const $window = $(window),
        $body = $('body'),
        $header = $('#header'),
        $banner = $('#banner');

    // Breakpoints.
    breakpoints({
        xlarge: '(max-width: 1680px)',
        large: '(max-width: 1280px)',
        medium: '(max-width: 980px)',
        small: '(max-width: 736px)',
        xsmall: '(max-width: 480px)'
    });

    // Play initial animations on page load.
    $window.on('load', function () {
        window.setTimeout(function () {
            $body.removeClass('is-preload');
        }, 100);
    });

    // Header.
    if ($banner.length > 0
        && $header.hasClass('alt')) {

        $window.on('resize', function () {
            $window.trigger('scroll');
        });

        $banner.scrollex({
            bottom: $header.outerHeight(),
            terminate: function () {
                $header.removeClass('alt');
            },
            enter: function () {
                $header.addClass('alt');
            },
            leave: function () {
                $header.removeClass('alt');
            }
        });

    }

    $('.scrolly').scrolly();

    // const $menu = $('#menu');
    // $menu._locked = false;
    //
    // $menu._lock = function () {
    //     if ($menu._locked)
    //         return false;
    //
    //     $menu._locked = true;
    //     window.setTimeout(function () {
    //         $menu._locked = false;
    //     }, 350);
    //
    //     return true;
    // };
    //
    // $menu._show = function () {
    //     if ($menu._lock())
    //         $body.addClass('is-menu-visible');
    //
    // };
    //
    // $menu._hide = function () {
    //     if ($menu._lock())
    //         $body.removeClass('is-menu-visible');
    //
    // };
    //
    // $menu._toggle = function () {
    //     if ($menu._lock())
    //         $body.toggleClass('is-menu-visible');
    //
    // };
    //
    // $menu
    //     .appendTo($body)
    //     .on('click', function (event) {
    //
    //         event.stopPropagation();
    //
    //         $menu._hide();
    //     })
    //     .find('.inner')
    //     .on('click', '.close', function (event) {
    //
    //         event.preventDefault();
    //         event.stopPropagation();
    //         event.stopImmediatePropagation();
    //
    //         $menu._hide();
    //     })
    //     .on('click', function (event) {
    //         event.stopPropagation();
    //     })
    //     .on('click', 'a', function (event) {
    //         const href = $(this).attr('href');
    //
    //         event.preventDefault();
    //         event.stopPropagation();
    //         $menu._hide();
    //
    //         window.setTimeout(function () {
    //             window.location.href = href;
    //         }, 350);
    //
    //     });
    //
    // $body
    //     .on('click', 'a[href="#menu"]', function (event) {
    //         event.stopPropagation();
    //         event.preventDefault();
    //         $menu._toggle();
    //     })
    //     .on('keydown', function (event) {
    //         if (event.keyCode === 27 /* ESC */) $menu._hide();
    //     });

})(jQuery);
