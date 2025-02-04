-- create database todo_list
use todo_list
-- ALTER TABLE todo
-- ADD COLUMN image_url VARCHAR(255); -- Cột để lưu đường dẫn hoặc tên tệp hình ảnh


-- CREATE TABLE user (
--     id INT AUTO_INCREMENT PRIMARY KEY,       -- ID duy nhất cho mỗi người dùng
--     username VARCHAR(50) NOT NULL,           -- Tên người dùng
--     email VARCHAR(100) NOT NULL UNIQUE,      -- Email duy nhất
--     password VARCHAR(255) NOT NULL,          -- Mật khẩu đã mã hóa
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Thời gian tạo
-- );
-- CREATE TABLE todo (
--     id INT AUTO_INCREMENT PRIMARY KEY,       -- ID duy nhất cho mỗi công việc
--     user_id INT NOT NULL,                    -- ID người dùng (khóa ngoại từ bảng user)
--     title VARCHAR(255) NOT NULL,             -- Tiêu đề công việc
--     description TEXT,                        -- Mô tả chi tiết
--     completed BOOLEAN DEFAULT FALSE,         -- Trạng thái hoàn thành
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Thời gian tạo
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Thời gian cập nhật

--     -- Định nghĩa khóa ngoại
--     CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
-- );
-- CREATE TABLE roles (
--     id INT AUTO_INCREMENT PRIMARY KEY,   -- ID duy nhất cho mỗi vai trò
--     name VARCHAR(50) NOT NULL UNIQUE,    -- Tên vai trò (admin, user, etc.)
--     description TEXT                     -- Mô tả vai trò
-- );
-- ALTER TABLE user 
-- ADD COLUMN role_id INT DEFAULT 2, -- Mặc định là vai trò "user" (giả sử ID của role "user" là 2)
-- ADD CONSTRAINT fk_role FOREIGN KEY (role_id) REFERENCES roles(id);

-- INSERT INTO user (username, email, password) VALUES 
-- ('john_doe', 'john@example.com', 'password123'), 
-- ('jane_doe', 'jane@example.com', 'admin123');

-- INSERT INTO todo (user_id, title, description, completed) VALUES 
-- (1, 'Buy groceries', 'Milk, Bread, Eggs', FALSE),
-- (1, 'Clean the house', 'Clean the living room and kitchen', TRUE),
-- (2, 'Read a book', 'Finish reading The Great Gatsby', FALSE);

-- INSERT INTO roles (name, description) VALUES
-- ('admin', 'Quản trị viên hệ thống'),
-- ('user', 'Người dùng thông thường'),
-- ('editor', 'Người chỉnh sửa nội dung');
