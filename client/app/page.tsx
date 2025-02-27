export default async function Page() {
  const data = await fetch('http://server:8080/')
  const posts = await data.json()
  console.log(posts)
  return (
    <ul>
      {posts.Connected}
    </ul>
  )
}