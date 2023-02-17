reacthandler
==

serve reactjs project

**note:**

 If mount react app in sub path, like `router.GET("/static/*any", rh.GinHandler())`, you need add prefix `/static/` to NewHandler, and change react router basename, add a config to package.json file, as follows:

app.js:

 ```js
 import { BrowserRouter } from "react-router-dom";

 function App() {
  return (
    <div>
        <BrowserRouter basename="/static">
        </BrowserRouter>
    </div>)}
 ```

package.json:

 ```json
 {
	  "homepage": "/static"
 }
 ```