package docs

import (
	"fmt"
	"net/http"
	
	"net/url"
	"github.com/swaggo/http-swagger/v2"
)

func SwaggerHandler(uri *url.URL) http.HandlerFunc {
	uriStr := fmt.Sprintf("%s/swagger/doc.json", uri.String())
	
	return httpSwagger.Handler(
		httpSwagger.URL(uriStr),
		httpSwagger.BeforeScript(`const UrlMutatorPlugin = (system) => ({
  rootInjects: {
    setScheme: (scheme) => {
      const jsonSpec = system.getState().toJSON().spec.json;
      const schemes = Array.isArray(scheme) ? scheme : [scheme];
      const newJsonSpec = Object.assign({}, jsonSpec, { schemes });

      return system.specActions.updateJsonSpec(newJsonSpec);
    },
    setHost: (host) => {
      const jsonSpec = system.getState().toJSON().spec.json;
      const newJsonSpec = Object.assign({}, jsonSpec, { host });

      return system.specActions.updateJsonSpec(newJsonSpec);
    },
    setBasePath: (basePath) => {
      const jsonSpec = system.getState().toJSON().spec.json;
      const newJsonSpec = Object.assign({}, jsonSpec, { basePath });

      return system.specActions.updateJsonSpec(newJsonSpec);
    }
  }
});`),
		httpSwagger.Plugins([]string{"UrlMutatorPlugin"}),
		httpSwagger.UIConfig(map[string]string{
			"onComplete": fmt.Sprintf(`() => {
    window.ui.setScheme('%s');
    window.ui.setHost('%s');
    window.ui.setBasePath('%s');
  }`, uri.Scheme, uri.Host, uri.Path),
		}),
	)
	
}