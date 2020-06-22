<template>
  <div>
    <PostPaging />
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'
import PostPaging from '~/components/PostPaging.vue'

@Component({
  components: { PostPaging },
})
export default class IndexPage extends Vue {}
</script>
