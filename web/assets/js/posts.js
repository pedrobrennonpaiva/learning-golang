$('#new-post-form').on('submit', createPost);
$(document).on('click', '.like-post', likePost);
$(document).on('click', '.unlike-post', unlikePost);

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
