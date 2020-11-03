const path = require('path');
const CleanWebpackPlugin = require('clean-webpack-plugin');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const WasmPackPlugin = require('@wasm-tool/wasm-pack-plugin');

function loadOutput(environment) {
    const filename = environment.production ? 'scripts/[name].[hash].min.js' : 'scripts/[name].js';
    const chunkFilename = environment.production ? '[name].[hash].min.js' : '[name].js';
    return {
        path: path.resolve(__dirname, 'dist'),
        filename,
        chunkFilename,
        publicPath: '/'
    };
};

function loadPlugins(environment) {

    let wasmPackArgs = '--no-typescript';
    if (environment.production) wasmPackArgs += ' --release';

    return [
        new HtmlWebpackPlugin({
            template: './src/index.html',
            chunks: ['app'],
            hash: true
        }),
        new WasmPackPlugin({
            crateDirectory: path.resolve(__dirname, `./src/hello_world`),
            extraArgs: wasmPackArgs
        }),
        new CleanWebpackPlugin()
    ];
};

module.exports = environment => {

    const mode = environment.production ? 'production' : 'development';
    const entry = { app: './src/app.js' };
    const output = loadOutput(environment);
    const plugins = loadPlugins(environment);

    //Base Config
    const config = { mode, entry, output, plugins };

    //Additional Production Config
    if (environment.production) {
        config.devtool = 'source-map';
    }

    //Additional Development Config
    else {
        config.devServer = {
            contentBase: path.join(__dirname, 'dist'),
            open: true,
            overlay: { errors: true, warnings: false },
            port: 5000,
            watchContentBase: true
        };
    }

    return config;
};
