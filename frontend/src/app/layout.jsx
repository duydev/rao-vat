import Footer from './components/footer';
import Header from './components/header';
import './globals.css';

export const metadata = {
  title: 'Rao vặt online',
  description: 'Một trang rao vặt online đơn giản'
};

export default function RootLayout({ children }) {
  return (
    <html lang="vi">
      <body>
        <Header />
        {children}
        <Footer />
      </body>
    </html>
  );
}
