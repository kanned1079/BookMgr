<script setup lang="ts">
import {h, onMounted, ref} from 'vue'
import type {FormInst, FormRules} from 'naive-ui'
import {NButton, NForm, NFormItem, NInput, NInputNumber, NSpace, NTag, useMessage} from 'naive-ui'
import instance from "@/axios";
// import * as console from "node:console";
// import * as console from "node:console";

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

let autoFillInText = ref<string>('')
let showDrawer = ref<boolean>(false)
let newBook = ref<Book>({
  name: '',
  publisher: '',
  year: 0,
  remark: '',
  author: '',
  isbn: '',
  price: 0.00,
  residue: 0,
  cover_url: '',
})

let editType = ref<'add' | 'edit'>('add')
let bookList = ref<Book[]>([])
let searchType = ref<string>('name')
let searchTarget = ref<string>('')
let searchSort = ref<string>('ASC')

const message = useMessage()
let animated = ref<boolean>(false)

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

const columns = [
  {
    title: '#',
    key: 'id',
    render(row: Book) {
      return h('p', {}, row.id);
    }
  },
  {
    title: '书名',
    key: 'name',
    render(row: Book) {
      return h('p', {}, row.name);
    }
  },
  {
    title: '出版社',
    key: 'publisher',
    render(row: Book) {
      return h('p', {}, row.publisher);
    }
  },
  {
    title: '发行年份',
    key: 'year',
    render(row: Book) {
      return h('p', {}, row.year);
    }
  },
  {
    title: '作者',
    key: 'author',
    render(row: Book) {
      return h('p', {}, row.author);
    }
  },
  {
    title: 'ISBN',
    key: 'isbn',
    render(row: Book) {
      return h('p', {}, row.isbn);
    }
  },
  {
    title: '价格',
    render(row: Book) {
      return h(NTag, {
        type: 'default',
        bordered: false,
      }, {default: () => row.price.toFixed(2)})
    }
  },
  {
    title: '剩余数目',
    render(row: Book) {
      return h(NTag, {
        type: 'default',
        bordered: false,
      }, {default: () => row.residue})
    }
  },
  {
    title: '操作',
    key: 'actions',
    render(row: Book) {
      return h('div', {style: {display: 'flex', flexDirection: 'row'}}, [
        h(NButton, {
          size: 'small',
          type: 'primary',
          bordered: false,
          style: {marginLeft: '10px'},
          onClick: () => editClicked(row),
        }, {default: () => '编辑'}),
        h(NButton, {
          size: 'small',
          type: 'error',
          style: {marginLeft: '10px'},
          onClick: () => deleteBookById(row),
        }, {default: () => '删除'})
      ]);
    }
  }
];

let editClicked = (row: Book) => {
  editType.value = 'edit'
  editBookId.value = row.id as number

  Object.keys(newBook.value).forEach(key => {
    // Cast `key` to `keyof typeof newBook.value` to let TypeScript know `key` is a valid key
    newBook.value[key as keyof typeof newBook.value] = row[key as keyof typeof row]
  })
  // Object.assign(editBookId.value, row)

  showDrawer.value = true
}

let getAllBooksAdmin = async () => {
  try {
    let {data} = await instance.get('/api/admin/v1/book', {
      params: {
        page: dataSize.value.page,
        size: dataSize.value.pageSize,
        search_by: searchType.value,
        search_content: searchTarget.value,
        search_sort: searchSort.value,
      }
    })
    if (data.code === 200) {
      bookList.value = []
      data.books.forEach((book: Book) => bookList.value.push(book))
      animated.value = true
      pageCount.value = data.page_count as number || 0
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

// -------

const formRef = ref<FormInst | null>(null)

const rules: FormRules = {
  name: [
    {required: true, message: '书名不能为空', trigger: 'blur'}
  ],
  publisher: [
    {required: true, message: '出版社不能为空', trigger: 'blur'}
  ],
  year: [
    {required: true, type: 'number', message: '发行年份不能为空', trigger: 'blur'},
    {validator: (rule, value) => value > 0, message: '发行年份应为有效的年份', trigger: 'blur'}
  ],
  author: [
    {required: true, message: '作者不能为空', trigger: 'blur'}
  ],
  isbn: [
    {required: true, message: 'ISBN不能为空', trigger: 'blur'}
  ],
  price: [
    {required: true, type: 'number', message: '价格不能为空', trigger: 'blur'},
    {validator: (rule, value) => value > 0, message: '价格应为正数', trigger: 'blur'}
  ],
  residue: [
    {required: true, type: 'number', message: '剩余藏书数量不能为空', trigger: 'blur'},
    {validator: (rule, value) => value >= 0, message: '剩余数量应为非负数', trigger: 'blur'}
  ],
  remark: [
    {required: true, message: '简述不能为空', trigger: 'blur'}
  ]
}

const handleCancel = () => {
  console.log("取消表单")
  clearForm()
  showDrawer.value = false
}

const handleAddNewBook = () => {
  formRef.value?.validate((errors) => {
    if (!errors) {
      submitNewBook()
      console.log("添加新书成功：", newBook.value)
    } else {
      console.log("表单验证失败", errors)
    }
  })
}

let submitNewBook = async () => {
  try {
    animated.value = false
    let {data} = await instance.post('/api/admin/v1/book', {
      ...newBook.value
    })
    if (data.code === 200) {
      showDrawer.value = false
      message.success('添加书目成功')
      await getAllBooksAdmin()
    } else {
      message.error('添加书目失败 ' + data.msg || '')
    }
  } catch (err: any) {
    message.error(err + '')
  }
}


let editBookId = ref<number>(-1)
let handleUpdateBookInfoById = async () => {
  try {
    animated.value = false
    let {data} = await instance.put('/api/admin/v1/book', {
      id: editBookId.value,
      ...newBook.value
    })
    if (data.code === 200) {
      showDrawer.value = false
      message.success('修改书目成功')
      await getAllBooksAdmin()
    } else {
      message.error('修改书目失败 ' + data.msg || '')
    }
  } catch (err: any) {
    message.error(err + '')
  }
}

let deleteBookById = async (row: Book) => {
  try {
    animated.value = false
    let {data} = await instance.delete('/api/admin/v1/book', {
      params: {
        id: row.id,
      }
    })
    if (data.code === 200) {
      showDrawer.value = false
      message.success('删除成功')
      await getAllBooksAdmin()
    } else {
      message.error('删除失败 ' + data.msg || '')
    }
  } catch (err: any) {
    message.error(err + '')
  }
}

// -------

let handleQuicklyFillIn = () => {
  const text = autoFillInText.value;
  newBook.value.name = text.match(/书名：([^，]+)/)?.[1] || '';
  newBook.value.publisher = text.match(/出版社：([^，]+)/)?.[1] || '';
  newBook.value.year = parseInt(text.match(/发行年份：(\d{4})/)?.[1] || '0', 10);
  newBook.value.author = text.match(/作者：([^，]+)/)?.[1] || '';
  newBook.value.isbn = text.match(/ISBN：([\dX]+)/)?.[1] || '';
  newBook.value.price = parseFloat(text.match(/价格：([\d.]+)/)?.[1] || '0');
  newBook.value.residue = parseInt(text.match(/剩余藏书数量：(\d+)/)?.[1] || '0', 10);
  newBook.value.remark = text.match(/简述：(.*)/)?.[1] || '';
};

let clearForm = () => {
  Object.keys(newBook.value).forEach(key => {
    if (typeof newBook.value[key] === 'string') {
      newBook.value[key] = '';
    } else if (typeof newBook.value[key] === 'number') {
      newBook.value[key] = 0;
    }
  });
};

onMounted(async () => {
  // message.info('BookMgr mounted.')
  await getAllBooksAdmin()
  animated.value = true
})

</script>

<script lang="ts">
export default {
  name: 'BookMgr'
}
</script>

<template>
  <div style="padding: 20px 20px 0 20px;">
    <n-card hoverable :embedded="true" :bordered="false">
      <div style="display: flex; flex-direction: row; justify-content: space-between; align-items: center">
        <p style="font-size: 1.1rem; font-weight: 500">图书管理</p>
        <n-button size="large" type="primary" @click="editType='add'; showDrawer=true">添加图书</n-button>
      </div>
    </n-card>
  </div>

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
          @click="animated=false; getAllBooksAdmin()"
      >
        搜索
      </n-button>
    </n-input-group>
  </div>

  <transition name="slide-fade">
    <div class="root" v-if="animated">

      <n-card :embedded="true" hoverable :bordered="false" content-style="padding: 0px" style="margin-top: 20px">
        <n-data-table
            class="table"
            :columns="columns"
            :data="bookList"
            :pagination="false"
            :bordered="true"
            style=""
            :scroll-x="800"
        />
      </n-card>

      <div style="margin-top: 20px; display: flex; flex-direction: row; justify-content: right;">
        <n-pagination
            size="medium"
            v-model:page.number="dataSize.page"
            :page-count="pageCount"
            @update:page="animated=false; getAllBooksAdmin() "
        />
        <n-select
            style="width: 160px; margin-left: 20px"
            v-model:value.number="dataSize.pageSize"
            size="small"
            :options="dataCountOptions"
            :remote="true"
            @update:value="animated=false; dataSize.page = 1; getAllBooksAdmin()"
        />
      </div>

    </div>
  </transition>

  <n-drawer
      v-model:show="showDrawer"
      :width="'60%'"
      :placement="'right'"
      @after-leave="clearForm()"
  >
    <div class="drawer-root">
      <p class="drawer-root-title">{{ editType === 'add' ? '添加新的图书信息' : `修改 ${newBook.name}` }}</p>

      <n-h4>快速输入</n-h4>
      <n-form-item label="图书信息 （回车确认）">
        <n-input @keyup.enter="handleQuicklyFillIn" type="textarea" v-model:value="autoFillInText" :rows="6" placeholder="书名：月亮与六便士，出版社：浙江文艺出版社，发行年份：2017，作者：毛姆，ISBN：9787533936020，价格：35.12，剩余藏书数量：17，简述：银行家查尔斯，人到中年，事业有成，为了追求内心隐秘的绘画梦想，突然抛妻别子，弃家出走。他深知：人的每一种身份都是一种自我绑架，唯有失去是通向自由之途。在异国他乡，他贫病交加，对梦想却愈发坚定执着。他说：我必须画画，就像溺水的人必须挣扎。"></n-input>
      </n-form-item>

      <n-h4>常规输入</n-h4>
      <div class="drawer-root-form">
        <n-form ref="formRef" :model="newBook" :rules="rules">
          <!-- 书名 -->
          <n-form-item path="name" label="书名">
            <n-input v-model:value="newBook.name" placeholder="请输入书名"/>
          </n-form-item>

          <!-- 出版社 -->
          <n-form-item path="publisher" label="出版社">
            <n-input v-model:value="newBook.publisher" placeholder="请输入出版社"/>
          </n-form-item>

          <!-- 发行年份 -->
          <n-form-item path="year" label="发行年份">
            <n-input-number style="width: 100%" v-model:value.number="newBook.year" placeholder="请输入发行年份"/>
          </n-form-item>

          <!-- 作者 -->
          <n-form-item path="author" label="作者">
            <n-input v-model:value="newBook.author" placeholder="请输入作者姓名"/>
          </n-form-item>

          <!-- ISBN -->
          <n-form-item path="isbn" label="ISBN">
            <n-input v-model:value="newBook.isbn" placeholder="请输入ISBN编号"/>
          </n-form-item>

          <!-- 价格 -->
          <n-form-item path="price" label="价格">
            <n-input-number style="width: 100%" v-model:value.number="newBook.price" placeholder="请输入价格"/>
          </n-form-item>

          <n-form-item path="residue" label="剩余藏书数量">
            <n-input-number style="width: 100%" v-model:value.number="newBook.residue" placeholder="请输入馆藏数量"/>
          </n-form-item>

          <n-form-item path="cover_url" label="图书封面">
            <n-input style="width: 100%" v-model:value="newBook.cover_url" placeholder="请输入封面图像链接"/>
          </n-form-item>

          <!-- 简述 -->
          <n-form-item path="remark" label="简述">
            <n-input v-model:value="newBook.remark" type="textarea" placeholder="请输入简述内容"/>
          </n-form-item>

          <!-- 底部按钮 -->
          <n-space align="center" justify="end" style="margin-top: 20px;">
            <n-button size="large" type="primary" tertiary @click="handleCancel">取消</n-button>
            <n-button size="large" v-if="editType==='add'" type="primary" @click="handleAddNewBook">确认添加</n-button>
            <n-button size="large" v-else type="primary" @click="animated=false; handleUpdateBookInfoById()">修改
            </n-button>
          </n-space>
        </n-form>
        <!--        </template>-->
      </div>
    </div>
  </n-drawer>

</template>

<style scoped lang="less">
.root {
  padding: 0 20px 20px 20px;
}

.drawer-root {
  padding: 30px;

  .drawer-root-title {
    font-size: 1.25rem;
    font-weight: 500;
  }

  .drawer-root-form {
    //margin-top: 20px;
  }
}
</style>