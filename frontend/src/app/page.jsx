import Image from 'next/image';
import Link from 'next/link';
import { getAllArticles } from '@/app/apis/articles-api';

async function Home() {
  const articles = await getAllArticles();

  console.log('articles', articles);

  return (
    <main className="max-w-[1140px] mx-auto my-4 grid grid-cols-4 gap-2">
      {articles.length === 0 && <div>Không có dữ liệu</div>}
      {articles.length > 0 &&
        articles.map(article => (
          <div key={article.id} className="">
            <Image
              className="w-full"
              src={`https://buffer.com/library/content/images/size/w1200/2023/10/free-images.jpg`}
              width={100}
              height={100}
              alt={article.title}
            />
            <Link className="text-lg font-bold" href={`/${article.id}`}>
              {article.title}
            </Link>
          </div>
        ))}
    </main>
  );
}

export default Home;
