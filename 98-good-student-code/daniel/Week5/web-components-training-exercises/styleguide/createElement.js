function createElement(elementName, templateSelector, proto) {
    proto = proto || HTMLElement.prototype;
    var currentScript = document._currentScript ? document._currentScript : document.currentScript;
    var template = currentScript.ownerDocument.querySelector(templateSelector).content;
    var customPrototype = Object.create(proto);
    var oldCallback = customPrototype.createdCallback;
    customPrototype.createdCallback = function () {
        var shadowRoot = this.createShadowRoot();
        var clone = document.importNode(template, true);
        shadowRoot.appendChild(clone);
        if (oldCallback) {
            oldCallback.call(this);
        }
    };
    return document.registerElement(elementName, {prototype: customPrototype});
}