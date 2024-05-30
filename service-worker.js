const CACHE_NAME = "pwa-cache-v1";

const HOSTNAME_WHITELIST = [
  self.location.hostname,
  "fonts.gstatic.com",
  "fonts.googleapis.com",
  "cdn.jsdelivr.net",
];

// The Util Function to hack URLs of intercepted requests
const getFixedUrl = (req) => {
  var now = Date.now();
  var url = new URL(req.url);

  // 1. fixed http URL
  // Just keep syncing with location.protocol
  // fetch(httpURL) belongs to active mixed content.
  // And fetch(httpRequest) is not supported yet.
  url.protocol = self.location.protocol;

  // 2. add query for caching-busting.
  // Github Pages served with Cache-Control: max-age=600
  // max-age on mutable content is error-prone, with SW life of bugs can even extend.
  // Until cache mode of Fetch API landed, we have to workaround cache-busting with query string.
  // Cache-Control-Bug: https://bugs.chromium.org/p/chromium/issues/detail?id=453190
  if (url.hostname === self.location.hostname) {
    url.search += (url.search ? "&" : "?") + "cache-bust=" + now;
  }
  return url.href;
};

self.addEventListener("install", (event) => {
  self.skipWaiting();
});

self.addEventListener("message", (event) => {
  if (event.data.action === "skipWaiting") {
    self.skipWaiting();
  }
});

/**
 *  @Lifecycle Activate
 *  New one activated when old isnt being used.
 *
 *  waitUntil(): activating ====> activated
 */
// Clean up old caches during the activate event
self.addEventListener("activate", (event) => {
  const cacheWhitelist = [CACHE_NAME];

  event.waitUntil(
    caches
      .keys()
      .then((cacheNames) => {
        return Promise.all(
          cacheNames.map((cacheName) => {
            if (!cacheWhitelist.includes(cacheName)) {
              return caches.delete(cacheName);
            }
          })
        );
      })
      .then(() => {
        return self.clients.claim();
      })
  );
});

self.addEventListener("fetch", (event) => {
  const url = new URL(event.request.url);

  // Define URLs or patterns to exclude from cache
  const shouldBypassCache =
    url.pathname.startsWith("/notifications") ||
    url.search.includes("no-cache") ||
    url.search.includes("realtime") ||
    // This is the path to accept invitations by token.
    url.pathname.startsWith("/i") ||
    // Bypass cache for navigations.
    event.request.mode === "navigate";

  // Bypass the service worker for non-GET requests.
  if (event.request.method !== "GET" || shouldBypassCache) {
    return;
  }
  event.respondWith(
    fetch(event.request).catch(() => {
      return caches.match(event.request);
    })
  );
});

self.addEventListener("push", (event) => {
  try {
    console.log("[Service Worker] Push Received.");

    const data = event.data.json();
    const title = data.title || "Chérie";
    const options = {
      body: data.body,
      icon: "https://chatbond-static.s3.us-west-002.backblazeb2.com/cherie/pwa/manifest-icon-96.maskable.png",
    };
    event.waitUntil(self.registration.showNotification(title, options));
  } catch (error) {
    console.error("Error in push event:", error);
  }
});

self.addEventListener("notificationclick", function (event) {
  // Close the notification when clicked
  event.notification.close();

  // Retrieve dynamic URL from the notification data
  const notificationData = event.notification.data;
  const targetUrl = notificationData
    ? notificationData.url
    : "/auth/notifications";

  // This looks to see if the current is already open and focuses if it is
  event.waitUntil(
    clients
      .matchAll({
        type: "window",
      })
      .then(function (clientList) {
        for (var i = 0; i < clientList.length; i++) {
          var client = clientList[i];
          if (client.url === targetUrl && "focus" in client) {
            return client.focus();
          }
        }
        if (clients.openWindow) {
          return clients.openWindow(targetUrl);
        }
      })
  );
});
