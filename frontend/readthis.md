# The README.md is for those that want to learn more about Next.js. For this project, connecting Next.js to Golang, do the following steps for Next.js:

1. First, run this to install the necessary packages
  - > npm install
Make sure the package_lock.json and node_modules are not present before running the command above as they can cause conflict

2. After the npm install, you can make the changes to your Next.js application.

3. Run your Next.js application:
  - > npm run start
  or
  - > npm run dev

4. Once the application is successful without any error or bug, build the application
  - > npm run build

5. After building, export the application so Golang can see it
  - > npm run export
  Note: The above command was included into the package.json file

6. Once export is done, go to the README.md for backend directory