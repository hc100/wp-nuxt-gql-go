<template>
  <div>
    <v-card>
      <v-container fluid grid-list-lg>
        <v-layout row wrap>
          <PostCard v-for="post in posts" :key="post.id" :post="post" />
        </v-layout>
      </v-container>
    </v-card>
  </div>
</template>
<script lang="ts">
import PostCard from '~/components/PostCard.vue'
import { Component, Vue } from '~/node_modules/nuxt-property-decorator'
import 'vue-apollo'
import postConnection from '~/apollo/queries/postConnection.gql'
import { Edge, EdgeOrder, PageCondition } from '~/gql-types'

class DataTableOptions {
  public page: number = 1
  public itemsPerPage: number = 3
  public sortBy: Array<string> = []
  public sortDesc: Array<boolean> = []
}

@Component({
  components: {
    PostCard,
  },
})
export default class PostPaging extends Vue {
  private readonly search = ''
  private posts = new Array<Node>()

  private totalCount: number = 0
  private startCursor: string | null = null
  private endCursor: string | null = null
  private nowPage: number = 1

  private options = new DataTableOptions()

  public created() {
    this.nowPage = 1
    this.options.page = 1
    this.connection()
  }

  private async connection() {
    try {
      const res = await this.$apollo.query({
        query: postConnection,
        variables: {
          filterWord: { filterWord: this.search },
          pageCondition: this.createPageCondition(
            this.nowPage,
            this.options.page,
            this.options.itemsPerPage,
            this.startCursor,
            this.endCursor
          ),
          edgeOrder: this.createEdgeOrder(
            this.options.sortBy,
            this.options.sortDesc
          ),
        },
      })

      if (res && res.data && res.data.postConnection) {
        const conn = res.data.postConnection

        this.posts = conn.edges
          .filter((e: Edge) => e.node)
          .map((e: Edge) => e.node)

        this.totalCount = conn.totalCount

        const pageInfo = conn.pageInfo
        this.startCursor = pageInfo.startCursor
        this.endCursor = pageInfo.endCursor

        this.nowPage = this.options.page
      } else {
        console.log('no result') // eslint-disable-line
      }
    } catch (e) {
      console.log(e) // eslint-disable-line
    }
  }

  private createPageCondition(
    nowPage: number,
    nextPage: number,
    limit: number,
    startCursor: string | null,
    endCursor: string | null
  ): PageCondition {
    return {
      forward: nowPage < nextPage ? { first: limit, after: endCursor } : null,
      backward:
        nowPage > nextPage ? { last: limit, before: startCursor } : null,
      nowPageNo: nowPage,
      initialLimit: limit > 0 ? limit : null,
    }
  }

  private createEdgeOrder(
    sortBy: Array<string>,
    sortDesc: Array<boolean>
  ): EdgeOrder | null {
    if (sortBy && sortDesc) {
      if (sortBy.length !== 1 || sortDesc.length !== 1) {
        return null
      }
      const direction = sortDesc[0] ? 'DESC' : 'ASC'
      return { key: { postOrderKey: 'POST_DATE' }, direction } as EdgeOrder
    }
    return null
  }
}
</script>
