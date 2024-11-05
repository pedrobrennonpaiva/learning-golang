$('#new-post-form').on('submit', createPost);

$(document).on('click', '.like-post', likePost);
$(document).on('click', '.unlike-post', unlikePost);

$('#update-post').on('click', updatePost);
$('.delete-post').on('click', deletePost);

function createPost(event) {
    event.preventDefault();

    const title = $('#title').val();
    const content = $('#content').val();

    $.ajax({
        url: '/posts',
        method: 'POST',
        data: {
            title,
            content
        }
    }).done(function() {
        window.location = '/';
    }).fail(function(e) {
        console.log(e);
        alert('Error creating post');
    });
}

function likePost(event) {
    event.preventDefault();

    const elementClicked = $(event.target);
    const postId = elementClicked.closest('div').data('post-id');

    elementClicked.prop('disabled', true);

    $.ajax({
        url: `/posts/${postId}/like`,
        method: 'POST'
    }).done(function() {
        const counter = elementClicked.next('span');
        const quantity = parseInt(counter.text());
        counter.text(quantity + 1);

        elementClicked.addClass('unlike-post');
        elementClicked.addClass('text-danger');
        elementClicked.removeClass('like-post');
    }).fail(function(e) {
        console.log(e);
        alert('Error liking post');
    }).always(function() {
        elementClicked.prop('disabled', false);
    });
}

function unlikePost(event) {
    event.preventDefault();

    const elementClicked = $(event.target);
    const postId = elementClicked.closest('div').data('post-id');

    elementClicked.prop('disabled', true);

    $.ajax({
        url: `/posts/${postId}/unlike`,
        method: 'POST'
    }).done(function() {
        const counter = elementClicked.next('span');
        const quantity = parseInt(counter.text());
        counter.text(quantity - 1);

        elementClicked.addClass('like-post');
        elementClicked.removeClass('unlike-post');
        elementClicked.removeClass('text-danger');
    }).fail(function(e) {
        console.log(e);
        alert('Error unliking post');
    }).always(function() {
        elementClicked.prop('disabled', false);
    });
}


function updatePost() {
    $(this).prop('disabled', true);

    const postId = $(this).data('post-id');
    
    $.ajax({
        url: `/posts/${postId}`,
        method: "PUT",
        data: {
            title: $('#title').val(),
            content: $('#content').val()
        }
    }).done(function() {
        Swal.fire('Success!', 'Post updated successfully!', 'success')
            .then(function() {
                window.location = "/";
            })
    }).fail(function() {
        Swal.fire("Ops...", "Error updating post!", "error");
    }).always(function() {
        $('#update-post').prop('disabled', false);
    })
}

function deletePost(event) {
    event.preventDefault();

    Swal.fire({
        title: "Warning!",
        text: "Are you sure you want to delete this post? This action is irreversible!",
        showCancelButton: true,
        cancelButtonText: "Cancel",
        icon: "warning"
    }).then(function(confirmation) {
        if (!confirmation.value) return;

        const elementClicked = $(event.target);
        const post = elementClicked.closest('div')
        const postId = post.data('post-id');
    
        elementClicked.prop('disabled', true);
    
        $.ajax({
            url: `/posts/${postId}`,
            method: "DELETE"
        }).done(function() {
            post.fadeOut("slow", function() {
                $(this).remove();
            });
        }).fail(function() {
            Swal.fire("Ops...", "Error deleting post!", "error");
        });
    })
}