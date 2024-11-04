$('#new-post-form').on('submit', createPost);
$('.like-post').on('click', likePost);

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
    console.log('likePost');

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
    }).fail(function(e) {
        console.log(e);
        alert('Error liking post');
    }).always(function() {
        elementClicked.prop('disabled', false);
    });
}