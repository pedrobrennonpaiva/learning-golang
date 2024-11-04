$('#register-form').on('submit', register);

function register(e) {
    e.preventDefault();

    const name = $('#name').val();
    const nickname = $('#nickname').val();
    const email = $('#email').val();
    const password = $('#password').val();
    const confirmPassword = $('#confirmPassword').val();

    if (password !== confirmPassword) {
        alert('As senhas não são iguais');
        return;
    }

    $.ajax({
        url: '/register',
        method: 'POST',
        data: {
            name,
            nickname,
            email,
            password
        }
    }).done(function() {
        alert('User registered successfully');
    }).fail(function(error) {
        console.log(error);
        alert('Error registering user');
    });
}
