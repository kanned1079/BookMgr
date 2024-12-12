<script setup lang="ts">
import {onMounted, ref} from 'vue'
import {useRouter} from "vue-router";
import instance from "@/axios";
import {NButton, NInput, useMessage} from "naive-ui";
import {ChevronForward, CheckmarkDone, CloseOutline} from '@vicons/ionicons5'
import {useUserStore} from "@/stores/userinfo";
import {formatDate} from "@/utils/timeFormat";
import useThemeStore from "@/stores/theme";
import theme from "@/stores/theme";

interface Book {
  id?: number
  name: string
  publisher: string
  year: number
  remark: string
  author: string
  isbn: string
  price: number
  residue: number
  cover_url?: string
  created_at?: string
  updated_at?: string
}

const themeStore = useThemeStore()
const userStore = useUserStore()
const router = useRouter()
let animated = ref<boolean>(false)
let showModal = ref<boolean>(false)
let bookList = ref<Book[]>([])
let bookDetail = ref<Book>({
  id: 0,
  name: '',
  publisher: '',
  year: 0,
  remark: '',
  author: '',
  isbn: '',
  price: 0.00,
  residue: 0,
  created_at: '',
  updated_at: '',
  cover_url: '',
})
let searchType = ref<string>('name')
let searchTarget = ref<string>('')
let searchSort = ref<string>('ASC')
const message = useMessage()

let pageCount = ref(10)

let dataSize = ref<{ pageSize: number, page: number }>({
  pageSize: 10,
  page: 1,
})

const bookDetails = (book: Book) => [
  {label: '作者', value: book.author},
  {label: '出版社', value: book.publisher},
  {label: '发行时间', value: book.year},
  {label: '价格', value: `${book.price.toFixed(2)} CNY`},
  {label: 'ISBN', value: book.isbn},
  {label: '入库时间', value: formatDate(book.created_at as string)},
  {label: '剩余藏书数量', value: book.residue}
];

let dataCountOptions = [
  {
    label: '10条数据/页',
    value: 10,
  },
  {
    label: '20条数据/页',
    value: 20,
  },
  {
    label: '50条数据/页',
    value: 50,
  },
  {
    label: '100条数据/页',
    value: 100,
  },
]

let selectFilter = [
  {
    label: '书名',
    value: 'name',
  },
  {
    label: '出版社',
    value: 'publisher',
  },
  {
    label: '发行年份',
    value: 'year',
  },
  {
    label: '作者',
    value: 'author',
  },
  {
    label: 'ISBN',
    value: 'isbn',
  },
  {
    label: '价格',
    value: 'price',
  }
]

let getAllBooks = async () => {
  try {
    let {data} = await instance.get('/api/user/v1/book', {
      params: {
        page: dataSize.value.page,
        size: 3,
        search_by: searchType.value,
        search_content: searchTarget.value,
        search_sort: searchSort.value,
      }
    })
    if (data.code === 200) {
      data.books.forEach((book: Book) => bookList.value.push(book))
      animated.value = true
      pageCount.value = data.page_count as number || 0
      if (data.books.length === 0) {
        message.warning('已经到底了')
      }
      if (bookList.value.length === 0) {
        message.warning('找不到符合条件的书目')
      } else {
        console.log('获取成功')
      }
    } else {
      message.error('获取失败 ' + data.msg || '')
    }
  } catch (err: any) {
    message.error(err + '')
  }
}

let borrowBookById = async (bookId: number) => {
  // message.info('bookId: ' + bookId)
  try {
    let {data} = await instance.post('/api/user/v1/borrow', {
      user_id: userStore.thisUser.id,
      book_id: bookId,
    })
    if (data.code === 200) {
      themeStore.menuSelected = 'my-books'
      await router.push({path: '/user/borrowed'})
      message.success('成功')
    } else if (data.code === 422) {
      message.error('剩余数量不足')
    } else {
      message.error('未知错误 ' + data.msg || '')
    }
  } catch (err: any) {
    message.error(err + '')
  }
}

let showMoreBooks = () => {
  animated.value = false
  dataSize.value.page += 1
  getAllBooks()
}

let showBookDetail = (book: Book) => {
  Object.assign(bookDetail.value, book)
  showModal.value = true
}

onMounted(async () => {
  await getAllBooks()
  animated.value = true
})
</script>

<script lang="ts">
export default {
  name: 'Borrow'
}
</script>

<template>
  <div style="padding: 20px 20px 0 20px; display: flex; flex-direction: row">
    <n-input-group>
      <n-select
          size="large"
          :style="{ flex: 2 }"
          :options="selectFilter"
          :default-value="'name'"
          v-model:value="searchType"
      />
      <n-select
          size="large"
          :style="{ flex: 1 }"
          :options="[{label: '升序', value: 'ASC'}, {label: '降序', value: 'DESC'}]"
          :default-value="'ASC'"
          v-model:value="searchSort"
      />
      <n-input
          placeholder="请输入查询条件（留空为查询所有）"
          size="large"
          :style="{ flex: 5 }"
          v-model:value="searchTarget"
      />
      <n-button
          size="large"
          type="primary"
          :style="{flex: 1}"
          @click="dataSize.page=1; bookList=[]; animated=false; getAllBooks()"
      >
        搜索
      </n-button>
    </n-input-group>
  </div>

  <transition name="slide-fade">
    <div class="root" v-if="animated">
      <n-grid cols="2 s:1 m:2 l:3" responsive="screen" x-gap="10px" y-gap="10px">
        <n-grid-item v-for="(book, k) in bookList" :key="book.id">
          <n-card hoverable :embedded="true" :bordered="false" :title="book.name">
            <template v-for="(detail, index) in bookDetails(book)" :key="index">
              <div class="book-item">
                <p class="book-item-title">{{ detail.label }}:</p>
                <p class="book-item-content">{{ detail.value }}</p>
              </div>
            </template>
            <div style="margin-top: 30px">
              <n-button type="primary" text @click="showBookDetail(book)">
                <p style="font-size: 1rem; margin-right: 5px">显示详情</p>
                <n-icon size="16">
                  <ChevronForward/>
                </n-icon>
              </n-button>
            </div>
          </n-card>
        </n-grid-item>
      </n-grid>
      <n-button type="primary" class="btn" @click="showMoreBooks">显示更多</n-button>
    </div>
  </transition>

  <n-modal v-model:show="showModal">
    <n-card
        style="width: 60%"
        title="详情"
        :bordered="false"
        size="huge"
        role="dialog"
        aria-modal="true"
    >
      <div class="detail-root">
        <div class="detail-root-left">
          <div class="book-item">
            <p class="book-item-title">作者:</p>
            <p class="book-item-content">{{ bookDetail.author }}</p>
          </div>
          <div class="book-item">
            <p class="book-item-title">出版社:</p>
            <p class="book-item-content">{{ bookDetail.publisher }}</p>
          </div>
          <div class="book-item">
            <p class="book-item-title">发行时间:</p>
            <p class="book-item-content">{{ bookDetail.year }}</p>
          </div>
          <div class="book-item">
            <p class="book-item-title">价格:</p>
            <p class="book-item-content">{{ bookDetail.price.toFixed(2) }} CNY</p>
          </div>
          <div class="book-item">
            <p class="book-item-title">ISBN:</p>
            <p class="book-item-content">{{ bookDetail.isbn }}</p>
          </div>
          <div class="book-item">
            <p class="book-item-title">入库时间:</p>
            <p class="book-item-content">{{ formatDate(bookDetail.created_at as string) }}</p>
          </div>
          <div class="book-item">
            <p class="book-item-title">剩余藏书数量:</p>
            <p class="book-item-content">{{ bookDetail.residue }}</p>
          </div>
          <div>
            <p style="font-size: 1.1rem; font-weight: 500;margin-right: 10px;opacity: 0.6;">描述:</p>
            <p style="font-size: 1rem;opacity: 0.9;">{{ bookDetail.remark }}</p>
          </div>

          <n-button
              :disabled="bookDetail.residue===0" @click="borrowBookById(bookDetail.id as number)"
              type="primary"
              style="margin-top: 40px; width: 60%; height: 40px"
          >
            {{ bookDetail.residue === 0 ? '剩余数量不足' : `借这个` }}
            <n-icon v-if="bookDetail.residue === 0" style="margin-left: 6px" size="20"><CloseOutline/></n-icon>
            <n-icon v-else style="margin-left: 6px" size="18"><CheckmarkDone/></n-icon>

          </n-button>
        </div>

        <div class="detail-root-right">
          <div class="img">
            <n-image class="img-inner" width="100%" :src="bookDetail.cover_url"></n-image>
          </div>
        </div>
      </div>

    </n-card>
  </n-modal>

</template>

<style scoped lang="less">
.root {
  padding: 20px;
}

.book-item {
  display: flex;
  flex-direction: row;
  margin-bottom: 10px;
  align-items: center;

  .book-item-title {
    font-size: 1.1rem;
    font-weight: 500;
    margin-right: 10px;
    opacity: 0.6;
  }

  .book-item-content {
    font-size: 1rem;
    opacity: 0.9;
  }
}

.btn {
  margin-top: 50px;
  width: 100%;
  height: 50px;
}

.detail-root {
  display: flex;
  flex-direction: row;

  .detail-root-left {
    margin-right: 60px;
    flex: 3;
    //background-color: skyblue;
  }

  //.detail-root-right {
  //  margin-left: 10px;
  //  flex: 2;
  //  display: flex;
  //  justify-content: center;
  //  align-items: center;
  //
  //  .img {
  //    width: 100%;
  //    height: 100%;
  //    background-size: contain; /* 确保图片不被裁剪 */
  //    background-position: center;
  //    background-repeat: no-repeat; /* 防止图片重复 */
  //
  //    .img-inner {
  //      box-shadow: 5px 5px 10px rgba(113, 113, 113, 0.8);
  //      border-radius: 3px;
  //      transition: ease 300ms;
  //      margin-bottom: 20px;
  //    }
  //
  //    .img-inner:hover {
  //      transform: translateX(0) translateY(-3px);
  //    }
  //  }
  //}

  .detail-root-right {
    margin-left: 10px;
    flex: 2;
    display: flex;
    justify-content: center;
    align-items: center;

    .img {
      width: 100%;
      height: 100%;
      background-size: cover; /* 将图片放大到填充整个容器 */
      background-position: center;
      background-repeat: no-repeat; /* 防止图片重复 */

      .img-inner {
        box-shadow: 5px 5px 10px rgba(113, 113, 113, 0.8);
        border-radius: 3px;
        transition: ease 300ms;
        margin-bottom: 20px;
      }

      .img-inner:hover {
        transform: translateX(0) translateY(-3px);
      }
    }
  }
}

</style>