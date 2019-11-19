import { sparky, fusebox, pluginLink } from 'fuse-box';
import * as path from 'path';

const { runDev, runProd } = fusebox({
  output: 'dist/server/$name-$hash',
  target: 'server',
  entry: 'server/server.ts',
  tsConfig: 'server/tsconfig.json',
  dependencies: { include: ['tslib'] },
  cache: {
    enabled: false,
    root: '.cache/server'
  }
  // codeSplitting: { scriptRoot: path.resolve(__dirname, './dist/server') }
});

(async () => {
  try {
    await runProd({ uglify: false, manifest: false });
  } catch (e) {
    console.error(e);
  }
})();
