var template = '' +
    '<div class="email-item email-item-selected pure-g">' +
        '<div class="pure-u">' +
            '<img class="email-avatar" alt="Yahoo! Finance Avatar" height="64" width="64" src="img/yfinance-avatar.png">' +
        '</div>' +

        '<div class="pure-u-3-4">' +
            '<h5 class="email-name">Yahoo! Finance</h5>' +
            '<h4 class="email-subject">How to protect your finances from winter storms</h4>' +
            '<p class="email-desc">' +
                'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Necessitatibus, sequi.' +
            '</p>' +
        '</div>' +
    '</div>';

    var result = '';
    for (var i = 0; i < 200; i++) {
      result += template;
    }

    document.querySelector('#list').innerHTML += result;
