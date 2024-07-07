import MagnifyingGlass from './MagnifyingGlass';
import './header.css';

function Header() {
  return (
    <header className="bg-[#ec3222] text-white">
      <div className="max-w-[1140px] m-auto">
        <div className="flex text-sm">
          <ul className="list-none flex">
            <li className="p-2 separator">
              <a href="#">Điều khoản sử dụng</a>
            </li>
            <li className="p-2 separator">
              <a href="#">Hướng dẫn đăng tin</a>
            </li>
            <li className="p-2 separator">
              <a href="#">Các câu hỏi thường gặp</a>
            </li>
          </ul>
          <ul className="list-none flex ml-auto">
            <li className="p-2 separator">
              <a href="#">Đăng ký</a>
            </li>
            <li className="p-2 separator">
              <a href="#">Đăng nhập</a>
            </li>
          </ul>
        </div>
        <div className="flex items-center">
          <div className="text-3xl">RAO VẶT</div>
          <div className="p-3 bg-[#d00000] mx-4 my-2 rounded-md grow flex">
            <input
              type="text"
              placeholder="Bạn muốn tìm gì?"
              className="bg-inherit text-md placeholder:text-white w-full"
            />
            <button className='ml-2'>
              <MagnifyingGlass size={20} />
            </button>
          </div>
          <button className="bg-[#ff9400] p-2 rounded-md">Đăng tin</button>
        </div>
      </div>
      <div className="bg-[#d00000]">
        <div className="max-w-[1140px] m-auto">
          <ul className="list-none flex justify-center text-lg">
            <li className="p-2 separator">
              <a href="#">Nhà đất</a>
            </li>
            <li className="p-2 separator">
              <a href="#">Xe</a>
            </li>
            <li className="p-2 separator">
              <a href="#">Học hành - Tuyển dụng</a>
            </li>
            <li className="p-2 separator">
              <a href="#">Sản phẩm dịch vụ</a>
            </li>
          </ul>
        </div>
      </div>
    </header>
  );
}

export default Header;
