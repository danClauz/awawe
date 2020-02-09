ALTER TABLE post_tag
    ADD CONSTRAINT FK_Posts
        FOREIGN KEY (post_id) REFERENCES posts (id);

ALTER TABLE post_tag
    ADD CONSTRAINT FK_Tags
        FOREIGN KEY (tag_id) REFERENCES tags (id);