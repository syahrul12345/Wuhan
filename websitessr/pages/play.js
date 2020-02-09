import { Cookies } from 'react-cookie';
import { handleAuthSSR } from '../src/utils/auth';
export default function Play() {
    // set up cookies
    const cookies = new Cookies();
    return(
        <div> PLAY PAGE </div>
    )
}

Play.getInitialProps = async(ctx) => {
    await handleAuthSSR(ctx);
    return {
        'status':true
    }
}