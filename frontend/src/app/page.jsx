import { getAllArticles } from '@/app/apis/articles-api';

async function Home() {
  const articles = await getAllArticles();

  console.log(articles);

  return <main className="max-w-[1140px] mx-auto my-4">Hello World</main>;
}

export default Home;
