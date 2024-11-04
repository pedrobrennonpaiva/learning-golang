$('#login-form').on('submit', login);

function login(e) {
    e.preventDefault();

    const email = $('#email').val();
    const password = $('#password').val();

    $.ajax({
        url: '/login',
        method: 'POST',
        data: {
            email,
            password
        }
    }).done(function() {
        window.location = '/home';
    }).fail(function(error) {
        console.log(error);
        alert('Error logging in');
    });
}
