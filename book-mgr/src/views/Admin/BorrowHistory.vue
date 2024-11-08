<script setup lang="ts">
import {ref, onMounted, h} from 'vue'
import {NButton, useMessage, NTag, NInput} from "naive-ui";
import instance from "@/axios";

const message = useMessage()

let animated = ref<boolean>(false)

interface SearchOptions {
  type: string
  target: string
}

let searchOptions = ref<SearchOptions>({
  type: 'email',
  target: '',
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
    label: '用户邮箱',
    value: 'email',
  },
  {
    label: '书名',
    value: 'name',
  },
  {
    label: 'ISBN',
    value: 'isbn',
  },
]

let pageCount = ref(10)

let dataSize = ref<{ pageSize: number, page: number }>({
  pageSize: 10,
  page: 1,
})

interface BorrowHistory {
  id: number
  borrow_id: string
  email: string
  book_name: string
  book_isbn: string
  created_at: string
  is_back: boolean
}

let histories = ref<BorrowHistory[]>([])

let getAllHistory = async () => {
  try {
    let {data} = await instance.get('/api/admin/v1/history', {
      params: {
        page: dataSize.value.page,
        size: dataSize.value.pageSize,
        search_type: searchOptions.value.type,  // 'email' 搜索用户邮箱   'name' 'isbn' 搜索书名或isbn号
        search_target: searchOptions.value.target.trim(), // 如果为空则搜索所有条目 按照created_at DESC搜索
      }
    })
    if (data.code === 200) {
      histories.value = []
      data.histories.forEach((history: BorrowHistory) => histories.value.push(history))
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
      return h('p', {}, row.id?.toString() || '-');
    }
  },
  {
    title: '借书订单号',
    key: 'borrow_id',
    render(row: BorrowHistory) {
      return h('p', {}, row.borrow_id || '-');
    }
  },
  {
    title: '用户邮箱',
    key: 'email',
    render(row: BorrowHistory) {
      return h('p', {}, row.email || '-');
    }
  },
  {
    title: '书名',
    key: 'book_name',
    render(row: BorrowHistory) {
      return h('p', {}, row.book_name || '-');
    }
  },
  {
    title: 'ISBN',
    key: 'book_isbn',
    render(row: BorrowHistory) {
      return h('p', {}, row.book_isbn || '-');
    }
  },
  {
    title: '借阅时间',
    key: 'created_at',
    render(row: BorrowHistory) {
      return h('p', {}, row.created_at);
    }
  },
  {
    title: '操作',
    key: 'actions',
    render(row: BorrowHistory) {
      return h('div', {style: {display: 'flex', flexDirection: 'row'}}, [
        h(NButton, {
          size: 'small',
          type: 'primary',
          secondary: !row.is_back,
          tertiary: row.is_back,
          bordered: false,
          disabled: row.is_back,
          style: {marginLeft: '1px', width: '100px'},
          onClick: () => message.info('功能还在开发'),
        }, {default: () => !row.is_back?'提醒还书':'已还'}),
      ]);
    }
  }
];


onMounted(async () => {
  await getAllHistory()
  animated.value = true
})

</script>

<script lang="ts">
export default {
  name: 'BorrowHistory'
}
</script>

<template>
  <div style="padding: 20px 20px 0 20px;">
    <n-card hoverable :embedded="true" :bordered="false">
      <div style="display: flex; flex-direction: row; justify-content: space-between; align-items: center">
        <p style="font-size: 1.1rem; font-weight: 500">借阅记录</p>
      </div>
    </n-card>
  </div>

  <div style="padding: 20px 20px 0 20px; display: flex; flex-direction: row">
    <n-input-group>
      <n-select
          size="large"
          :style="{ flex: 2 }"
          :options="selectFilter"
          :default-value="'email'"
          v-model:value="searchOptions.type"
      />
<!--      <n-select-->
<!--          size="large"-->
<!--          :style="{ flex: 1 }"-->
<!--          :options="[{label: '升序', value: 'ASC'}, {label: '降序', value: 'DESC'}]"-->
<!--          :default-value="'ASC'"-->
<!--          v-model:value="searchOptions.sort"-->
<!--      />-->
      <n-input
          placeholder="请输入查询条件（留空为默认查询所有符合条件的行）"
          size="large"
          :style="{ flex: 5 }"
          v-model:value="searchOptions.target"
      />
      <n-button
          size="large"
          type="primary"
          :style="{flex: 1}"
          @click="animated=false; getAllHistory()"
      >
        搜索
      </n-button>
    </n-input-group>
  </div>

  <transition name="slide-fade">
    <div class="root" v-if="animated">
      <n-card :embedded="true" hoverable :bordered="false" content-style="padding: 0px" style="margin-top: 20px">
        <n-data-table
            striped
            class="table"
            :columns="columns"
            :data="histories"
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
            @update:page="animated=false; getAllHistory() "
        />
        <n-select
            style="width: 160px; margin-left: 20px"
            v-model:value.number="dataSize.pageSize"
            size="small"
            :options="dataCountOptions"
            :remote="true"
            @update:value="animated=false; dataSize.page = 1; getAllHistory()"
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