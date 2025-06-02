// // import type { NextConfig } from "next";

// // const nextConfig: NextConfig = {
// //   /* config options here */
// // };

// // export default nextConfig;

// // /** /** @type {import('next').NextConfig} */
// const nextConfig = {
//   reactStrictMode: true,
// };

// const withPWA = require("next-pwa")({
//   dest: "public", // PWA files destination
//   register: true, // Register SW automatically
//   skipWaiting: true, // Skip waiting for service worker to activate
//   disable: process.env.NODE_ENV === "development", // Disable PWA in dev mode
// });

// module.exports = withPWA(nextConfig);

/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
};

const withPWA = require("next-pwa")({
  dest: "public",
  register: true,
  skipWaiting: true,
  disable: process.env.NODE_ENV === "development",
});

module.exports = withPWA(nextConfig);