class BlocksView extends HTMLElement {
  connectedCallback() {
    this.resourceConfig().then(config => {
      this.fetchComponents(() => {
        this.bootComponents(config);
      });
    });
  }

  disconnectedCallback() {}

  /**
   * Fetches the resource configuration from the server. The source is loaded from the `resource`
   * attribute in the element tag.
   *
   * @return {Promise<any>}
   */
  resourceConfig() {
    return fetch(this.url(`resource_config?id=${this.resourceName()}`)).then(
      res => res.json()
    );
  }

  /**
   * Sends an HTTP request to fetch the JS code for the asked components.
   */
  fetchComponents(callback) {
    let script = document.createElement("script");
    script.src = "http://127.0.0.1:4005/dist/main.js";
    script.onload = callback;

    document.head.appendChild(script);
  }

  /**
   * Initialize components
   */
  bootComponents(props) {
    if (this.getAttribute("load") === "crud") {
      Blocks.ReactDOM.render(Blocks.Crud(props), this);
    }
  }

  /**
   * Formats a full URL to send a request to the blocks server.
   *
   * @param {string} path Path of the URL
   * @return {string}
   */
  url(path) {
    return "http://127.0.0.1:8080/" + path;
  }

  /**
   * Returns the name of the resource for this `blocks-view` tag.
   *
   * @return {string}
   */
  resourceName() {
    return this.getAttribute("resource");
  }
}

window.customElements.define("blocks-view", BlocksView);
