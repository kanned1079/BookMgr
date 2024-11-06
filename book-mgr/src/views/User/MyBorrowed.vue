<script setup lang="ts">
import {onMounted, ref, h} from 'vue'
import instance from "@/axios";
import {useUserStore} from "@/stores/userinfo";
import {useMessage, NTag, NButton} from "naive-ui";

const userStore = useUserStore()
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

interface BorrowHistory {
  id: number  // 数据库id
  borrow_id: string // 生成的借书订单号
  created_at: string  // 借书日期
  is_back: boolean  // 是否归还
  keep: string  // 留存时间 (当前时间 - created_at 在后端完成)
  name: string  // 书名
  isbn: string  // 书ISBN
}

let history = ref<BorrowHistory[]>([])
let searchBookName = ref<string>('')

let getAllMyHistory = async () => {
  try {
    let {data} = await instance.get('/api/user/v1/history', {
      params: {
        user_id: userStore.thisUser.id,
        page: dataSize.value.page,
        size: dataSize.value.pageSize,
        name: searchBookName.value.trim(),  // 如果为空则查询所有
      }
    })
    if (data.code === 200) {
      history.value = []
      data.histories.forEach((i: BorrowHistory) => history.value.push(i))
      pageCount.value = data.page_count
      animated.value = true
    }
  } catch (err: any) {
    message.error(err + '')
  }
}

const columns = [
  {
    title: 'ID',
    key: 'id',
    render(row: BorrowHistory) {
      return h('p', {}, row.id.toString());  // 显示数据库 ID
    }
  },
  {
    title: '借书订单号',
    key: 'borrow_id',
    render(row: BorrowHistory) {
      return h('p', {}, row.borrow_id);  // 显示借书订单号
    }
  },
  {
    title: '书名',
    key: 'name',
    render(row: BorrowHistory) {
      return h('p', {}, row.name);  // 显示书名
    }
  },
  {
    title: 'ISBN',
    key: 'isbn',
    render(row: BorrowHistory) {
      return h('p', {}, row.isbn);  // 显示书的 ISBN
    }
  },
  {
    title: '借书日期',
    key: 'created_at',
    render(row: BorrowHistory) {
      return h('p', {}, row.created_at);  // 显示借书日期
    }
  },
  {
    title: '留存时间',
    key: 'keep',
    render(row: BorrowHistory) {
      return h('p', {}, row.keep);  // 显示留存时间
    }
  },
  {
    title: '是否归还',
    key: 'is_back',
    render(row: BorrowHistory) {
      return h(
          NTag,
          {
            type: row.is_back ? 'success' : 'error', // 已归还为 success，未归还为 error
            bordered: false,
          },
          { default: () => (row.is_back ? '已归还' : '未归还') }  // 根据状态显示标签文本
      );
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
        }, {default: () => '还书'}),
        // h(NButton, {
        //   size: 'small',
        //   type: 'error',
        //   style: {marginLeft: '10px'},
        //   onClick: () => deleteBookById(row),
        // }, {default: () => '删除'})
      ]);
    }
  }
];

onMounted(async () => {
  await getAllMyHistory()
  animated.value = true
})

</script>

<script lang="ts">
export default {
  name: 'MyBorrowed'
}
</script>

<template>
  <div style="padding: 20px 20px 0 20px;">
    <n-card hoverable :embedded="true" :bordered="false">
      <div style="display: flex; flex-direction: row; justify-content: space-between; align-items: center">
        <p style="font-size: 1.1rem; font-weight: 500">我借的书</p>
        <!--        <n-button size="large" type="primary" @click="editType='add'; showDrawer=true">添加图书</n-button>-->
      </div>
    </n-card>
  </div>

  <transition name="slide-fade">
    <div class="root" v-if="animated">
      <n-card :embedded="true" hoverable :bordered="false" content-style="padding: 0px" style="margin-top: 20px">
        <n-data-table
            class="table"
            :columns="columns"
            :data="history"
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
            @update:page="animated=false; getAllMyHistory() "
        />
        <n-select
            style="width: 160px; margin-left: 20px"
            v-model:value.number="dataSize.pageSize"
            size="small"
            :options="dataCountOptions"
            :remote="true"
            @update:value="animated=false; dataSize.page = 1; getAllMyHistory()"
        />
      </div>

    </div>
  </transition>

</template>

<style scoped lang="less">
.root {
  padding: 0 20px 20px 20px;
}
</style>