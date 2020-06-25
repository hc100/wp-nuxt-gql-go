<template>
  <v-row>
    <div>
      <div class="headline">{{ post.post_title }}</div>
      <div>
        <span v-if="post.category">カテゴリ: {{ post.category.name }}</span>
        <span v-if="post.tags"
          >タグ:
          <span v-for="(list, index) in post.tags">
            <span>{{ list.name }}</span
            ><span v-if="index + 1 < post.tags.length">, </span>
          </span>
        </span>
      </div>
      <div>{{ post.post_date }}</div>
      <div>
        <pre class="wp_content" v-html="post.post_content" />
      </div>
    </div>
  </v-row>
</template>

<script lang="ts">
import { Component, Prop, Vue } from '~/node_modules/nuxt-property-decorator'

@Component
export default class PostCard extends Vue {
  @Prop({ default: null })
  post!: Object
}
</script>

<style>
pre {
  white-space: pre-wrap; /* Since CSS 2.1 */
  white-space: -moz-pre-wrap; /* Mozilla, since 1999 */
  white-space: -pre-wrap; /* Opera 4-6 */
  white-space: -o-pre-wrap; /* Opera 7 */
  word-wrap: break-word; /* Internet Explorer 5.5+ */
  margin-bottom: 20px;
}
</style>
