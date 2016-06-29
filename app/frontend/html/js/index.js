document.addEventListener("WebComponentsReady", function() {
	var googleSigninButton = document.querySelector("#google-signin-button");

	googleSigninButton.addEventListener("google-signin-success", function(e) {
		idToken = gapi.auth2.getAuthInstance().currentUser.get().getAuthResponse().id_token;

		var xhr = new XMLHttpRequest();
		xhr.open("POST", "/v1/auth/google/verify");
		xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
		
		xhr.onload = function() {
			console.log("/v1/auth/google/verify", xhr.responseText);
		};

		xhr.send("id_token=" + idToken);
	});
});
