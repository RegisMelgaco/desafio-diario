<script type="ts">
  import { Router, Route } from "svelte-routing";
  import Login from "./pages/Login.svelte";

  import Cookies from 'js-cookie'

  export let url = "";

  if (!Cookies.get("credential")) {
    url = "/login"
    console.log(url)
  }
</script>

<Router {url}>
  <Route path="/login"><Login /></Route>
</Router>
