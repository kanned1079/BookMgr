<script setup lang="ts">
import {onMounted, ref} from 'vue'
import instance from "@/axios";
import {NButton, NInput, useMessage} from "naive-ui";
import {ChevronForward} from '@vicons/ionicons5'
import {formatDate} from "@/utils/timeFormat";

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
  created_at?: string
  updated_at?: string
}

let animated = ref<boolean>(false)

let bookList = ref<Book[]>([])
let searchType = ref<string>('name')
let searchTarget = ref<string>('')
let searchSort = ref<string>('ASC')

const message = useMessage()

let pageCount = ref(10)

let dataSize = ref<{ pageSize: number, page: number }>({
  pageSize: 10,
  page: 1,
})

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
      // bookList.value = []
      data.books.forEach((book: Book) => bookList.value.push(book))
      animated.value = true
      pageCount.value = data.page_count as number || 0
      if (data.books.length === 0) {
        message.warning('已经到底了')
      }
      if (bookList.value.length === 0) {
        message.warning('找不到符合条件的书目')
      } else {
        // message.success('获取成功')
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
  message.info('bookId: ' + bookId)
}

let showMoreBooks = () => {
  animated.value = false
  dataSize.value.page += 1
  getAllBooks()
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
          placeholder="请输入查询条件（留空为默认查询所有符合条件的行）"
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
      <n-grid cols="2 s:2" responsive="screen" x-gap="10px" y-gap="10px">
        <n-grid-item v-for="(book, k) in bookList" :key="book.id">
          <n-card hoverable :embedded="true" :bordered="false" :title="book.name">
            <div class="book-item">
              <p class="book-item-title">作者:</p>
              <p class="book-item-content">{{ book.author }}</p>
            </div>
            <div class="book-item">
              <p class="book-item-title">出版社:</p>
              <p class="book-item-content">{{ book.publisher }}</p>
            </div>
            <div class="book-item">
              <p class="book-item-title">发行时间:</p>
              <p class="book-item-content">{{ book.year }}</p>
            </div>
            <div class="book-item">
              <p class="book-item-title">价格:</p>
              <p class="book-item-content">{{ book.price.toFixed(2) }} CNY</p>
            </div>
            <div class="book-item">
              <p class="book-item-title">ISBN:</p>
              <p class="book-item-content">{{ book.isbn }}</p>
            </div>
            <div class="book-item">
              <p class="book-item-title">入库时间:</p>
              <p class="book-item-content">{{ formatDate(book.created_at) }}</p>
            </div>
            <div class="book-item">
              <p class="book-item-title">剩余藏书数量:</p>
              <p class="book-item-content">{{ book.residue }}</p>
            </div>
            <div style="margin-top: 30px">
              <n-button type="primary" text>
                <p style="font-size: 1rem; margin-right: 5px">显示详情</p>
                <n-icon size="16"><ChevronForward/></n-icon>
              </n-button>
            </div>

            <n-button @click="borrowBookById(book.id)" secondary style="margin-top: 30px; width: 40%" type="primary">借这个</n-button>
          </n-card>
        </n-grid-item>
      </n-grid>

      <n-button type="primary" class="btn" @click="showMoreBooks">显示更多</n-button>

    </div>
  </transition>

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

</style>