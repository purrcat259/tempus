const { FuseBox } = require('fuse-box');

const fuse = FuseBox.init({
  homeDir: 'src',
  output: 'dist/$name.js'
});

fuse.dev({ port: 4445, httpServer: false });

fuse
  .bundle('server/bundle')
  .watch('server/**')
  .instructions(' > [server/index.ts]')
  .completed(proc => proc.start());

// fuse.bundle("client/app")
//     .watch("client/**")
//     .hmr()
//     .instructions(" > client/index.ts");

fuse.run();
