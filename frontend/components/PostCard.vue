<template>
  <v-flex xs12>
    <v-card color="black">
      <v-container fluid grid-list-lg>
        <v-layout row>
          <v-flex xs7>
            <div>
              <div class="headline">{{ post.post_title }}</div>
              <div>{{ post.post_date }}</div>
              <div v-html="post.post_content" />
            </div>
          </v-flex>
        </v-layout>
      </v-container>
    </v-card>
  </v-flex>
</template>

<script>
export default {
  props: {
    post: {
      type: Object,
      default: null,
    },
  },
}
</script>
