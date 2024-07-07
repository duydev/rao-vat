/** @type {import('next').NextConfig} */
const nextConfig = {
  images: {
    remotePatterns: [
      {
        protocol: 'https',
        hostname: 'buffer.com'
      }
    ]
  }
};

export default nextConfig;
