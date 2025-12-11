import { Bot } from 'grammy';
import commands from './cmd/index';
import { Context } from 'grammy';

const bot = new Bot('');

for (const { cmd, handler } of commands) {
    bot.command(cmd, handler);
}

bot.command('help', (ctx: Context) => {
    console.log(ctx.message);

    let result = '';

    for (const { cmd, description } of commands) {
        result += `/${cmd}: ${description}\n`;
    }

    ctx.reply(result);
});

bot.start();
