<svelte:head>
	<title>Login</title>
	<meta name="description" content="Login page" />
</svelte:head>

<script>
  import {redirect} from '@sveltejs/kit'
  export let form

  if (form) {
    redirect(303, '/')
  }
</script>

<div class="container">
  <div class="row justify-content-center">

    <div class="col-lg-7">

      <div class="card o-hidden border-0 shadow-lg my-5">
        <div class="card-body p-0">
          <div class="row">
            <div class="col-lg">
              <div class="p-5">
                <div class="text-center">
                  <h1 class="h4 text-gray-900 mb-4">Sign in</h1>
                </div>

                  <form method="POST">
                    <div class="form-group">
                      <input type="text" class="form-control form-control-user" id="email" name="email" aria-describedby="emailHelp" placeholder="Enter Email Address...">
                    </div>

                    <div class="form-group">
                      <input type="password" class="form-control form-control-user" id="password" name="password" placeholder="Password">
                    </div>

                    <button type="submit" class="btn btn-primary">Login</button>
                  </form>

                <div class="text-center">
                  <a class="small" href="/register">Create an Account!</a>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

    </div>

  </div>

</div>