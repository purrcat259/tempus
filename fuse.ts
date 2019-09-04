import { sparky, fusebox, pluginLink } from 'fuse-box';
import * as path from 'path';
class Context {
  isProduction;
  runServer;
  getServerConfig() {
    return fusebox({
      output: 'dist/server/$name-$hash',
      target: 'server',
      entry: 'server/server.ts',
      tsConfig: 'server/tsconfig.json',
      dependencies: { include: ['tslib'] },
      cache: {
        enabled: false,
        root: '.cache/server'
      },
      codeSplitting: { scriptRoot: path.resolve(__dirname, './dist/server') }
    });
  }
  getBrowserConfig() {
    return fusebox({
      output: 'dist/browser/$name-$hash',
      target: 'browser',
      entry: 'src/index.tsx',
      tsConfig: 'src/tsconfig.json',
      dependencies: { include: ['tslib'] },
      webIndex: {
        // publicPath: '/public',
        template: 'src/index.html'
      },
      link: { useDefault: true },

      cache: {
        enabled: false,
        root: '.cache/browser'
      },
      watch: {
        chokidar: { usePolling: true },
        ignored: [/\.gql$/]
      },
      devServer: {
        //   httpServer: true,
        //   hmrServer: { port: 7878 }
        proxy: [
          {
            path: '/gql',
            options: {
              target: 'http://localhost:8000',
              changeOrigin: true,
              followRedirects: true,
              pathRewrite: {
                '^/gql': '/'
              }
            }
          }
        ]
      }
    });
  }
}
const { task, exec, rm } = sparky<Context>(Context);

task('default', async ctx => {
  await rm('./dist');
  const browser = ctx.getBrowserConfig();
  await browser.runDev();

  const server = ctx.getServerConfig();
  await server.runDev();
});

task('preview', async ctx => {
  ctx.isProduction = true;
  await rm('./dist');

  const browser = ctx.getBrowserConfig();
  await browser.runProd({ uglify: true, manifest: true });
  const server = ctx.getServerConfig();
  await server.runProd({ uglify: false, manifest: true });

  // response.launcher.start();
});
task('dist', async ctx => {
  ctx.runServer = false;
  ctx.isProduction = true;
  const fuse = ctx.getBrowserConfig();
  await fuse.runProd({ uglify: false, manifest: true });
});
