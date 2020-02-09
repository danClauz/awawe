ALTER TABLE comments
    ADD CONSTRAINT FK_Users
        FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE comments
    ADD CONSTRAINT FK_Posts
        FOREIGN KEY (post_id) REFERENCES posts (id);