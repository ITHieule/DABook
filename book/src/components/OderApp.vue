<template>
    <div class="order-page">
        <h1>Danh Sách Sách</h1>

        <!-- Danh sách sách -->
        <div class="book-list">
            <div v-for="book in books" :key="book.ID" class="book-item">
                <h3>{{ book.Ten }}</h3>
                <p>Số lượng còn: {{ book.Soluong }}</p>
                <button :disabled="book.Soluong === 0" @click="addToCart(book)">
                    Thêm vào giỏ hàng
                </button>
            </div>
        </div>

        <!-- Giỏ hàng -->
        <h2>Giỏ Hàng</h2>
        <div v-if="cart.length > 0" class="cart">
            <div v-for="item in cart" :key="item.ID" class="cart-item">
                <p>Số lượng (tối đa: {{ books.find(b => b.ID === item.ID)?.Soluong || 0 }}):</p>
                <input type="number" v-model="item.Soluong" min="1" @input="validateInput(item)" />


                <p>Số lượng: {{ item.Soluong }} </p>

                <button @click="removeFromCart(item)">Xóa</button>
                <button @click="confirmOrder(item)">Xác nhận mua</button>
            </div>
        </div>
        <p v-else>Giỏ hàng trống.</p>
    </div>
</template>

<script>
import axios from 'axios';

export default {
    data() {
        return {
            Soluong: 1, // Giá trị mặc định
            books: [],  // Danh sách sách
            cart: [],   // Giỏ hàng
        };
    },

    methods: {
        // Lấy danh sách sách từ API
        async fetchBooks() {
            try {
                const response = await axios.get('http://localhost:8081/api/v1/book/getbook');
                this.books = response.data.data;
            } catch (error) {
                console.error('Lỗi khi lấy danh sách sách:', error);
            }
        },
        // Thêm sách vào giỏ hàng
        addToCart(book) {
            const cartItem = this.cart.find((item) => item.ID === book.ID);
            if (cartItem) {
                if (cartItem.Soluong < book.Soluong) {
                    cartItem.Soluong++;
                } else {
                    alert(`Số lượng đã đạt tối đa: ${book.Soluong}`);
                }
            } else {
                this.cart.push({ ...book, Soluong: 1 });
            }
        },

        // Xóa sách khỏi giỏ hàng
        removeFromCart(item) {
            const index = this.cart.indexOf(item);
            if (index > -1) {
                this.cart.splice(index, 1);
            }
        },
        // Xác nhận mua và giảm số lượng sách trong kho
        async confirmOrder(item) {
            try {
                // Định dạng ngày đặt hàng (yyyy/MM/dd)
                const today = new Date();
                const formattedDate = today.toLocaleDateString('en-GB').split('/').reverse().join('/'); // Định dạng DD/MM/YYYY thành YYYY/MM/DD
                console.log(item);

                // Gửi yêu cầu xác nhận mua
                const response = await axios.post('http://localhost:8081/api/v1/book/Oderbook', {
                    IDkhachhang: Number(localStorage.getItem("idkhachhang")),
                    IDbook: item.ID,
                    Soluong: item.Soluong,
                    ngaydathang: formattedDate,
                });

                if (response.data.code == 200) {
                    alert('Đặt hàng thành công!');
                    this.cart = [];  // Xóa giỏ hàng
                    this.fetchBooks();  // Làm mới danh sách sách
                } else {
                    alert('Đặt hàng thất bại: ' + response.data.message);
                }
            } catch (error) {
                console.error('Lỗi khi xác nhận mua:', error);
            }
        },
        validateInput(item) {
            const bookInStock = this.books.find((book) => book.ID === item.ID); // Lấy sách tương ứng trong danh sách
            if (item.Soluong < 1) {
                item.Soluong = 1; // Không cho phép số lượng nhỏ hơn 1
            } else if (bookInStock && item.Soluong > bookInStock.Soluong) {
                item.Soluong = bookInStock.Soluong; // Giới hạn số lượng không vượt quá số lượng trong kho
                alert(`Số lượng không được vượt quá ${bookInStock.Soluong}!`);
            }
        },

    },
    created() {
        this.fetchBooks(); // Lấy danh sách sách khi component được tạo
    },
};
</script>

<style scoped>
.order-page {
    max-width: 800px;
    margin: 0 auto;
    padding: 20px;
}

.book-list {
    display: flex;
    flex-wrap: wrap;
    gap: 20px;
}

.book-item {
    border: 1px solid #ccc;
    padding: 10px;
    border-radius: 5px;
    width: 200px;
}

button {
    margin-top: 10px;
    padding: 5px 10px;
    background-color: #4caf50;
    color: white;
    border: none;
    cursor: pointer;
}

button:disabled {
    background-color: #ccc;
}

.cart {
    margin-top: 20px;
    border-top: 2px solid #4caf50;
    padding-top: 10px;
}

.cart-item {
    margin-bottom: 10px;
}

p {
    font-size: 16px;
}
</style>