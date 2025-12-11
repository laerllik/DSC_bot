import { Context } from 'grammy';

const handler = (ctx: Context) => {
    ctx.reply('pong');
};

export default {
    cmd: 'ping',
    handler,
    description: 'простой пинг бота',
};
