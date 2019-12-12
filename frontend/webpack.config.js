const webpack = require('webpack');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const TerserPlugin = require('terser-webpack-plugin');
var path = require('path');

module.exports = {
    entry: path.join(__dirname, 'src', 'app.js'),
    output: {
        path: path.join(__dirname, 'dist'),
        filename: 'bundle.js'
    },

    devServer: {
        contentBase: path.join(__dirname, 'dist'),
        // compress: true,
        port: 9000
    },
    module: {
        rules: [
            {
                test: /\.js$/,
                loader: 'babel-loader'
            },
            {
                test: /\.css$/,
                use: [
                    'style-loader',
                    'css-loader'
                ]
            },
            {
                test: /\.png$/,
                use: "url-loader?limit=100000"
            },
            {
                test: /\.jpg$/,
                use: "file-loader"
            },
            {
                test: /\.(woff|woff2)(\?v=\d+\.\d+\.\d+)?$/,
                use: 'url-loader?limit=10000&mimetype=application/font-woff'
            },
            {
                test: /\.ttf(\?v=\d+\.\d+\.\d+)?$/,
                use: 'url-loader?limit=10000&mimetype=application/octet-stream'
            },
            {
                test: /\.eot(\?v=\d+\.\d+\.\d+)?$/,
                use: 'file-loader'
            },
            {
                test: /\.svg(\?v=\d+\.\d+\.\d+)?$/,
                use: 'url-loader?limit=10000&mimetype=image/svg+xml'
            }
        ]
    },
    optimization: {
        minimize: true,
        minimizer: [new TerserPlugin({
            parallel: true,
            terserOptions: {
                output: {
                    comments: false,
                },
            },
            extractComments: false,
        })],
    },
    plugins: [
        new HtmlWebpackPlugin({
            template: path.resolve(__dirname, 'public', 'index.html'),
            filename: 'index.html'
        }),
        new webpack.NamedModulesPlugin(),
        new webpack.HotModuleReplacementPlugin(),
        new webpack.DefinePlugin({
            WEBSOCKET_URL: JSON.stringify(process.env.WEBSOCKET_URL || 'ws://localhost:8081/ws'),
            BACKEND_URL: JSON.stringify(process.env.BACKEND_URL || 'http://localhost:8081/api'),
            MAP_SERVER_URL: JSON.stringify(process.env.MAP_SERVER_URL || 'https://{a-c}.tile.openstreetmap.org/{z}/{x}/{y}.png'),
            MAP_PROJ: JSON.stringify(process.env.MAP_PROJ || 'EPSG:3857'),
            MAP_ZOOM: JSON.stringify(process.env.MAP_ZOOM || 2),
            MAP_CENTER: process.env.MAP_CENTER || [0, 0],
        })
    ]
};