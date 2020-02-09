ALTER TABLE post_category
    ADD CONSTRAINT FK_Posts
        FOREIGN KEY (post_id) REFERENCES posts (id);

ALTER TABLE post_category
    ADD CONSTRAINT FK_Categories
        FOREIGN KEY (category_id) REFERENCES categories (id);