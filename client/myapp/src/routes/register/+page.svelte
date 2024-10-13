<svelte:head>
	<title>Login</title>
	<meta name="description" content="Login page" />
</svelte:head>

<script>
  export let form
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
                      <input type="text" class="form-control form-control-user" id="name" name="name" aria-describedby="emailHelp" placeholder="Enter username">
                    </div>

                    <div class="form-group">
                      <input type="text" class="form-control form-control-user" id="email" name="email" aria-describedby="emailHelp" placeholder="Enter Email Address">
                    </div>

                    <div class="form-group">
                      <input type="password" class="form-control form-control-user" id="password" name="password" placeholder="Password">
                    </div>

                    <div class="form-group">
                      <input type="password" class="form-control form-control-user" id="password_confirmation" name="password_confirmation" placeholder="Password Confirmation">
                    </div>

                    <button type="submit" class="btn btn-primary">Register</button>
                  </form>

                  {#if form?.error}
                    <p class="alert alert-danger">{form.error}</p>
                  {/if}

                <div class="text-center">
                  <a class="small" href="/login">Already have an Account!</a>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

    </div>

  </div>

</div>
