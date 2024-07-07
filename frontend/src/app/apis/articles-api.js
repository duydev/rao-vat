import request from './request';

export async function getAllArticles() {
  const { data } = await request.get('/api/articles');

  return data;
}
