import type { ReactNode } from 'react';
import clsx from 'clsx';
import Link from '@docusaurus/Link';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import Layout from '@theme/Layout';
import Heading from '@theme/Heading';

import styles from './index.module.css';

function HomepageHeader() {
  const { siteConfig } = useDocusaurusContext();
  return (
    <header className={clsx('hero hero--primary', styles.heroBanner)}>
      <div className="container">
        <Heading as="h1" className="hero__title">
          {siteConfig.title}
        </Heading>
        <p className="hero__subtitle">{siteConfig.tagline}</p>
        <div className={styles.buttons}>
          <Link
            className="button button--secondary button--lg"
            to="/docs/getting-started"
          >
            Get Started
          </Link>
          <Link
            className="button button--outline button--lg"
            href="https://github.com/MikhailWahib/Ra"
          >
            GitHub
          </Link>
        </div>
      </div>
    </header>
  );
}

function HomepageMain(): ReactNode {
  return (
  <div className={styles.mainContent}>
      <div className="container">
        <div className={styles.codeExample}>
          <h2>Example Code</h2>
          <pre>
            <code>
              {`let hello = fn(name) {
  puts("Hello, " + name + "!");
}

hello("World");`}
            </code>
          </pre>
        </div>

        <div className={styles.features}>
          <h2>Why Ra?</h2>
          <ul>
            <li>Simple and intuitive syntax</li>
            <li>Built-in support for arrays and dictionaries</li>
            <li>Designed for learning and experimentation</li>
          </ul>
        </div>

        <div className={styles.cta}>
          <h2>Ready to Get Started?</h2>
          <Link
            className="button button--primary button--lg"
            to="/docs/getting-started"
          >
            Explore the Docs
          </Link>
        </div>
      </div>
    </div>
  );
}

export default function Home(): ReactNode {
  const { siteConfig } = useDocusaurusContext();
  return (
    <Layout
      title="Ra Programming Language"
      description="A simple interpreted language built in Go"
    >
      <HomepageHeader />
      <main>
        <HomepageMain />
      </main>
    </Layout>
  );
}