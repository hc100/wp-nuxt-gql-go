export default (error, context) => {
  console.log(error) // eslint-disable-line
  context.error({ statusCode: 304, message: 'Server error' })
}
