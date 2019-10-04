class Blocks extends HTMLElement {
  constructor() {
    super()
    this.resource = this.getAttribute('resource')
    this.load = this.getAttribute('load')
  }
}

window.customElements.define('blocks', Blocks)
