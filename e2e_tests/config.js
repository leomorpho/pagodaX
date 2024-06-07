const config = {
  development: {
    WEBSITE_URL: "http://localhost:8001",
    EMAIL_URL: "http://localhost:8020",
  },
  production: {
    WEBSITE_URL: "https://yourproductionurl.com",
  },
  // other environments...
};

// Default configuration
const defaultConfig = {
  WEBSITE_URL: "http://localhost:8001",
};

module.exports = { config, defaultConfig };
