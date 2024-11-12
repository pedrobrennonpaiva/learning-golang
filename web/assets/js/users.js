$('#register-form').on('submit', register);
$('#follow').on('click', followUser);
$('#unfollow').on('click', unfollowUser);
$('#edit-user-form').on('submit', editUser);
$('#change-password-form').on('submit', changePassword);
$('#delete-user').on('click', deleteUser);

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
        Swal.fire('Success!', 'User registered successfully', 'success');
    }).fail(function(error) {
        console.log(error);
        Swal.fire('Error!', 'Error registering user', 'error');
    });
}

function followUser(e) {
    e.preventDefault();

    const userId = $(this).data('user-id');

    $(this).prop('disabled', true);

    $.ajax({
        url: `/users/${userId}/follow`,
        method: 'POST'
    }).done(function() {
        window.location = `/users/${userId}`;
    }).fail(function(error) {
        console.log(error);
        Swal.fire('Error!', 'Error following user', 'error');
        $(this).prop('disabled', false);
    });
}

function unfollowUser(e) {
    e.preventDefault();
    
    const userId = $(this).data('user-id');

    $(this).prop('disabled', true);

    $.ajax({
        url: `/users/${userId}/unfollow`,
        method: 'POST'
    }).done(function() {
        window.location = `/users/${userId}`;
    }).fail(function(error) {
        console.log(error);
        Swal.fire('Error!', 'Error unfollowing user', 'error');
        $(this).prop('disabled', false);
    });
}

function editUser(e) {
    e.preventDefault();

    const name = $('#name-user').val();
    const nickname = $('#nickname').val();
    const email = $('#email').val();

    $.ajax({
        url: '/edit-user',
        method: 'PUT',
        data: {
            name,
            nickname,
            email
        }
    }).done(function() {
        Swal.fire('Success!', 'User updated successfully', 'success')
            .then(function() {
                window.location = '/profile';
            });
    }).fail(function(error) {
        console.log(error);
        Swal.fire('Error!', 'Error updating user', 'error');
    });
}

function changePassword(e) {
    e.preventDefault();

    const currentPassword = $('#current-password').val();
    const newPassword = $('#new-password').val();
    const confirmPassword = $('#confirm-password').val();

    if (newPassword !== confirmPassword) {
        Swal.fire('Ops...', 'The passwords are not the same', 'warning');
        return;
    }

    $.ajax({
        url: '/change-password',
        method: 'POST',
        data: {
            currentPassword,
            newPassword
        }
    }).done(function() {
        Swal.fire('Success!', 'Password updated successfully', 'success')
            .then(function() {
                window.location = '/profile';
            });
    }).fail(function(error) {
        console.log(error);
        Swal.fire('Error!', 'Error updating password', 'error');
    });
}

function deleteUser() {
    Swal.fire({
        title: 'Are you sure?',
        text: "You won't be able to revert this!",
        icon: 'warning',
        showCancelButton: true,
        cancelButtonText: 'Cancel',
        confirmButtonText: 'Yes, delete it!'
    }).then((result) => {
        if (result.isConfirmed) {
            $.ajax({
                url: '/delete-user',
                method: 'DELETE'
            }).done(function() {
                Swal.fire('Success!', 'User deleted successfully', 'success')
                    .then(function() {
                        window.location = '/logout';
                    });
            }).fail(function(error) {
                console.log(error);
                Swal.fire('Error!', 'Error deleting user', 'error');
            });
        }
    });
}