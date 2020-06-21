<template>
  <div>
    <v-row>
      <v-col cols="12" sm="6" offset-sm="3">
        <v-card>
          <v-list two-line subheader>
            <v-list-item v-for="post in posts" :key="post.id" link>
              <v-list-item-avatar>
                <v-icon>mdi-gift-outline</v-icon>
              </v-list-item-avatar>
              <v-list-item-content>
                <v-list-item-title>{{ post.post_title }}</v-list-item-title>
                <v-list-item-subtitle>{{
                  post.post_date
                }}</v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>
          </v-list>
        </v-card>
      </v-col>
    </v-row>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'
import 'vue-apollo'
import posts from '~/apollo/queries/posts.gql'
import { Post } from '~/gql-types'

@Component({
  apollo: {
    posts: {
      prefetch: true,
      query: posts,
    },
  },
})
export default class PostCard extends Vue {
  posts: Post[] = []
}
</script>
