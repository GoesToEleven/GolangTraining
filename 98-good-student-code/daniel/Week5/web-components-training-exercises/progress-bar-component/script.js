(function() {
  var progressBarShadowRoot = document.getElementById('progress-bar').createShadowRoot();
  progressBarShadowRoot.innerHTML += '<style>.total {max-width: 100%; border: 1px solid rgba(39,41,45,.20); border-radius: .25em; background: rgba(211, 211, 211, 0.24); padding: .25em; height: 1.8em;} .bar {background-color: #8de674; height: 1.75em;}</style>';
  progressBarShadowRoot.innerHTML += '<div class="total"><div class="bar" style="width: 45%;"></div></div><div id="label"></div>';
  var labelShadowRoot = progressBarShadowRoot.getElementById('label').createShadowRoot();
  labelShadowRoot.innerHTML += '<style>.label { width: 100%; text-align: center; }</style>';
  labelShadowRoot.innerHTML += '<div class="label">45% Complete</div>';
}());