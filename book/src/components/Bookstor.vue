<template>
  <div class="book-list">
    <h1>Danh sách Sách</h1>
    <!-- Hiển thị bảng sách -->
    <table v-if="books.length > 0" class="book-table">
      <thead>
        <tr>
          <th>ID</th>
          <th>Tên Sách</th>
          <th>ID Loại</th>
          <th>Số Lượng</th>
          <th>Thao Tác</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="book in books" :key="book.ID">
          <td>{{ book.ID }}</td>
          <td>{{ book.Ten }}</td>
          <td>{{ book.IDloai }}</td>
          <td>{{ book.Soluong }}</td>
          <td>
            <button @click="deleteBook(book.ID)" class="delete-button">Xóa</button>
            <button @click="editBook(book)" class="update-button">Sửa</button>
          </td>
        </tr>
      </tbody>
    </table>
    <p v-else>Không có sách nào được tìm thấy!</p>
    <button @click="fetchBooks" class="refresh-button">Làm mới</button>
    <!-- Form cập nhật sách -->
    <div>
      <h2>Cập nhật Sách</h2>
      <form @submit.prevent="updateBook(selectedBook)">
        <input type="hidden" v-model="selectedBook.ID" />
        <input type="text" v-model="selectedBook.Ten" placeholder="Tên Sách" required />
        <input type="number" v-model="selectedBook.IDloai" placeholder="ID Loại" required min="1"/>
        <input type="number" v-model="selectedBook.Soluong" placeholder="Số Lượng" required min="0" />
        <button type="submit" class="add-button">Cập nhật</button>
      </form>
    </div>
    <!-- Form thêm sách mới -->
    <div>
      <h2>Thêm Sách</h2>
      <form @submit.prevent="addBook">
        <input type="text" v-model="newBook.Ten" placeholder="Tên Sách" required />
        <input type="number" v-model="newBook.IDloai" placeholder="ID Loại" required min="1" />
        <input type="number" v-model="newBook.Soluong" placeholder="Số Lượng" required min="0"/>
        <button type="submit" class="add-button">Thêm</button>
      </form>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "BookstoreList",
  data() {
    return {
      books: [], // Danh sách sách
      newBook: { Ten: "", IDloai: "", Soluong: "" }, // Thông tin sách mới
      selectedBook: { ID: "", Ten: "", IDloai: "", Soluong: 0 }, // Thông tin sách đang được cập nhật
    };
  },
  mounted() {
    this.fetchBooks(); // Gọi API để lấy danh sách sách khi component được tải
  },
  methods: {
    // Lấy danh sách sách
    async fetchBooks() {
      try {
        const response = await axios.get("http://localhost:8081/api/v1/book/getbook");
        this.books = response.data.data; // Lấy danh sách sách
      } catch (error) {
        console.error("Lỗi khi gọi API:", error);
      }
    },
    // Thêm sách mới
    async addBook() {
      try {
        const response = await axios.post("http://localhost:8081/api/v1/book/Addbook", this.newBook);
        if (response.data.code === 200) {
          this.fetchBooks(); // Làm mới danh sách sau khi thêm sách
          this.newBook = { Ten: "", IDloai: "", Soluong: "" }; // Reset form
        } else {
          alert("Thêm sách không thành công!");
        }
      } catch (error) {
        console.error("Lỗi khi thêm sách:", error);
      }
    },
    // Xóa sách
    async deleteBook(bookID) {
      try {
        // Sending the request with content-type 'application/json' (Axios handles this automatically, but this can be explicitly set)
        const response = await axios.post(
          "http://localhost:8081/api/v1/book/Deletebook",

          { ID: String(bookID) },
          // Sending the bookID as JSON in the body
          {
            headers: {
              "Content-Type": "application/json", // Ensuring the request content type is set to JSON
            },
          }
        );

        if (response.data.code === 200) {
          this.fetchBooks(); // Refresh the book list after successful deletion
        } else {
          alert("Xóa sách không thành công!"); // Show alert if deletion fails
        }
      } catch (error) {
        console.error("Lỗi khi xóa sách:", error); // Log any errors during the delete request
      }
    },

    // Cập nhật sách
    async updateBook() {
      // Ensure that ID is converted to string if it's a number
      const updatedBook = {
        ID: String(this.selectedBook.ID), // Convert ID to string
        ten: this.selectedBook.Ten,
        idloai: this.selectedBook.IDloai,
        soluong: this.selectedBook.Soluong,
      };

      try {
        const response = await axios.post("http://localhost:8081/api/v1/book/Updatebook", updatedBook);
        if (response.data.code === 200) {
          this.fetchBooks(); // Làm mới danh sách sau khi cập nhật
          this.selectedBook = { ID: "", ten: "", idloai: "", soluong: "" }; // Reset form
        } else {
          alert("Cập nhật sách không thành công!");
        }
      } catch (error) {
        console.error("Lỗi khi cập nhật sách:", error);
      }
    },

    // Chọn sách để cập nhật
    editBook(book) {
      this.selectedBook = { ...book }; // Sao chép thông tin sách vào selectedBook
    },

  },

};
</script>

<style scoped>
.book-list {
  max-width: 800px;
  margin: 20px auto;
  text-align: center;
}

.book-table {
  width: 100%;
  border-collapse: collapse;
  margin: 20px 0;
}

.book-table th,
.book-table td {
  border: 1px solid #ddd;
  padding: 8px;
  text-align: center;
}

.book-table th {
  background-color: #f4f4f4;
  color: #333;
}

.refresh-button,
.add-button,
.delete-button,
.update-button {
  background-color: #42b983;
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 4px;
  cursor: pointer;
}

.refresh-button:hover,
.add-button:hover,
.delete-button:hover,
.update-button:hover {
  background-color: #369d73;
}
</style>